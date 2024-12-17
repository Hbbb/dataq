package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"go.quinn.io/dataq/config"
	"go.quinn.io/dataq/pluginutil"
	"go.quinn.io/dataq/queue"
	"go.quinn.io/dataq/schema"
)

func (w *Worker) startPluginHandlers(ctx context.Context, messages chan Message) (map[string]chan *schema.PluginRequest, map[string]map[string]*queue.Task, error) {
	pluginReqs := make(map[string]chan *schema.PluginRequest)
	taskmap := make(map[string]map[string]*queue.Task)

	for id, cfg := range w.plugins {
		reqs := make(chan *schema.PluginRequest)
		pluginReqs[id] = reqs
		taskmap[id] = make(map[string]*queue.Task)

		go w.runPlugin(ctx, id, cfg, reqs, taskmap[id], messages)
	}

	return pluginReqs, taskmap, nil
}

func (w *Worker) runPlugin(ctx context.Context, id string, cfg *config.Plugin, reqs chan *schema.PluginRequest, taskmap map[string]*queue.Task, messages chan Message) {
	resps, err := pluginutil.Execute(ctx, cfg, reqs)
	if err != nil {
		sendErrorf(messages, "Failed to start plugin %s: %v", id, err)
		return
	}

	w.handlePluginResponses(ctx, id, resps, taskmap, messages)
}

func (w *Worker) handlePluginResponses(ctx context.Context, pluginID string, resps <-chan *schema.PluginResponse, taskmap map[string]*queue.Task, messages chan Message) {
	for resp := range resps {
		if resp.Closed {
			// The plugin closed, no more requests will be handled
			return
		}

		task := taskmap[resp.RequestId]

		// Handle errors from the plugin
		if resp.Error != "" {
			w.taskError(ctx, task, messages, fmt.Errorf("plugin %s: %s", pluginID, resp.Error))
			continue
		}

		// Handle item responses
		if resp.Item != nil {
			if err := w.handlePluginResponseItem(ctx, resp.Item, messages); err != nil {
				w.taskError(ctx, task, messages, err)
				continue
			}
		}

		// Handle action responses
		if resp.Action != nil {
			if err := w.handlePluginResponseAction(ctx, resp, messages); err != nil {
				w.taskError(ctx, task, messages, err)
				continue
			}
		}

		// Handle completion
		if resp.Done {
			w.handlePluginResponseDone(ctx, task, messages)
		}
	}
}

func (w *Worker) handlePluginResponseItem(ctx context.Context, item *schema.DataItem, messages chan Message) error {
	jsonBytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

	r := bytes.NewReader(jsonBytes)
	if _, err := w.cas.Store(ctx, r); err != nil {
		return err
	}

	sendInfo(messages, fmt.Sprintf("[kind: %s] [data: %s] [hash: %s]", item.Meta.Kind, item.Meta.Id, item.Meta.Hash))
	return nil
}

func (w *Worker) handlePluginResponseAction(ctx context.Context, resp *schema.PluginResponse, messages chan Message) error {
	newTask := queue.NewTask(*w.plugins[resp.PluginId], resp.Action)
	if err := w.queue.Push(ctx, newTask); err != nil {
		return err
	}
	sendInfo(messages, fmt.Sprintf("[task: %s] [plugin: %s] [action: %s]", newTask.Uid, newTask.PluginId, resp.Action.Name))
	return nil
}

func (w *Worker) handlePluginResponseDone(ctx context.Context, task *queue.Task, messages chan Message) {
	w.updateTaskState(ctx, task, queue.TaskStatusComplete, "", messages)
	sendDone(messages)
}
