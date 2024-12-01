// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: proto/dataq.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Plugin represents a data extraction plugin
type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Config      map[string]string `protobuf:"bytes,4,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	mi := &file_proto_dataq_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dataq_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_proto_dataq_proto_rawDescGZIP(), []int{0}
}

func (x *Plugin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Plugin) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type DataItemMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId    string `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Kind        string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Timestamp   int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ContentType string `protobuf:"bytes,5,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *DataItemMetadata) Reset() {
	*x = DataItemMetadata{}
	mi := &file_proto_dataq_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataItemMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataItemMetadata) ProtoMessage() {}

func (x *DataItemMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dataq_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataItemMetadata.ProtoReflect.Descriptor instead.
func (*DataItemMetadata) Descriptor() ([]byte, []int) {
	return file_proto_dataq_proto_rawDescGZIP(), []int{1}
}

func (x *DataItemMetadata) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *DataItemMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataItemMetadata) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *DataItemMetadata) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *DataItemMetadata) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

// DataItem represents a single piece of data extracted by a plugin
type DataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta    *DataItemMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	RawData []byte            `protobuf:"bytes,6,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
}

func (x *DataItem) Reset() {
	*x = DataItem{}
	mi := &file_proto_dataq_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataItem) ProtoMessage() {}

func (x *DataItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dataq_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataItem.ProtoReflect.Descriptor instead.
func (*DataItem) Descriptor() ([]byte, []int) {
	return file_proto_dataq_proto_rawDescGZIP(), []int{2}
}

func (x *DataItem) GetMeta() *DataItemMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *DataItem) GetRawData() []byte {
	if x != nil {
		return x.RawData
	}
	return nil
}

// PluginConfig represents the configuration for all plugins
type PluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugins []*Plugin `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty"`
}

func (x *PluginConfig) Reset() {
	*x = PluginConfig{}
	mi := &file_proto_dataq_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfig) ProtoMessage() {}

func (x *PluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dataq_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfig.ProtoReflect.Descriptor instead.
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return file_proto_dataq_proto_rawDescGZIP(), []int{3}
}

func (x *PluginConfig) GetPlugins() []*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

// PluginRequest represents a request to a plugin
type PluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId  string            `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	Config    map[string]string `protobuf:"bytes,2,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Item      *DataItem         `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	Operation string            `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"` // e.g., "extract", "configure"
}

func (x *PluginRequest) Reset() {
	*x = PluginRequest{}
	mi := &file_proto_dataq_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRequest) ProtoMessage() {}

func (x *PluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dataq_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRequest.ProtoReflect.Descriptor instead.
func (*PluginRequest) Descriptor() ([]byte, []int) {
	return file_proto_dataq_proto_rawDescGZIP(), []int{4}
}

func (x *PluginRequest) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *PluginRequest) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *PluginRequest) GetItem() *DataItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *PluginRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

// PluginResponse represents a response from a plugin
type PluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId string    `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	Item     *DataItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Error    string    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"` // empty if no error
}

func (x *PluginResponse) Reset() {
	*x = PluginResponse{}
	mi := &file_proto_dataq_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginResponse) ProtoMessage() {}

func (x *PluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dataq_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginResponse.ProtoReflect.Descriptor instead.
func (*PluginResponse) Descriptor() ([]byte, []int) {
	return file_proto_dataq_proto_rawDescGZIP(), []int{5}
}

func (x *PluginResponse) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *PluginResponse) GetItem() *DataItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *PluginResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_dataq_proto protoreflect.FileDescriptor

var file_proto_dataq_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x61, 0x74, 0x61, 0x71, 0x22, 0xbc, 0x01, 0x0a, 0x06, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x71, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x39,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x94, 0x01, 0x0a, 0x10, 0x44, 0x61,
	0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x52, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2b, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x71, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x61, 0x77,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x22, 0xe4, 0x01,
	0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x71, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x68, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x19,
	0x5a, 0x17, 0x67, 0x6f, 0x2e, 0x71, 0x75, 0x69, 0x6e, 0x6e, 0x2e, 0x69, 0x6f, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x71, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_dataq_proto_rawDescOnce sync.Once
	file_proto_dataq_proto_rawDescData = file_proto_dataq_proto_rawDesc
)

func file_proto_dataq_proto_rawDescGZIP() []byte {
	file_proto_dataq_proto_rawDescOnce.Do(func() {
		file_proto_dataq_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dataq_proto_rawDescData)
	})
	return file_proto_dataq_proto_rawDescData
}

var file_proto_dataq_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_dataq_proto_goTypes = []any{
	(*Plugin)(nil),           // 0: dataq.Plugin
	(*DataItemMetadata)(nil), // 1: dataq.DataItemMetadata
	(*DataItem)(nil),         // 2: dataq.DataItem
	(*PluginConfig)(nil),     // 3: dataq.PluginConfig
	(*PluginRequest)(nil),    // 4: dataq.PluginRequest
	(*PluginResponse)(nil),   // 5: dataq.PluginResponse
	nil,                      // 6: dataq.Plugin.ConfigEntry
	nil,                      // 7: dataq.PluginRequest.ConfigEntry
}
var file_proto_dataq_proto_depIdxs = []int32{
	6, // 0: dataq.Plugin.config:type_name -> dataq.Plugin.ConfigEntry
	1, // 1: dataq.DataItem.meta:type_name -> dataq.DataItemMetadata
	0, // 2: dataq.PluginConfig.plugins:type_name -> dataq.Plugin
	7, // 3: dataq.PluginRequest.config:type_name -> dataq.PluginRequest.ConfigEntry
	2, // 4: dataq.PluginRequest.item:type_name -> dataq.DataItem
	2, // 5: dataq.PluginResponse.item:type_name -> dataq.DataItem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_dataq_proto_init() }
func file_proto_dataq_proto_init() {
	if File_proto_dataq_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_dataq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_dataq_proto_goTypes,
		DependencyIndexes: file_proto_dataq_proto_depIdxs,
		MessageInfos:      file_proto_dataq_proto_msgTypes,
	}.Build()
	File_proto_dataq_proto = out.File
	file_proto_dataq_proto_rawDesc = nil
	file_proto_dataq_proto_goTypes = nil
	file_proto_dataq_proto_depIdxs = nil
}
