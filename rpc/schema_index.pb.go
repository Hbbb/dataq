// Code generated by protoc-gen-dataq-index. DO NOT EDIT.

package rpc

func (m *Email) SchemaKind() string {
	return "Email"
}

func (m *Email) SchemaMetadata() map[string]interface{} {
	metadata := make(map[string]interface{})

	if m.From != "" {
		metadata["from"] = m.From
	}
	if m.To != "" {
		metadata["to"] = m.To
	}
	if m.Subject != "" {
		metadata["subject"] = m.Subject
	}
	if m.Body != "" {
		metadata["body"] = m.Body
	}
	if m.Date != "" {
		metadata["date"] = m.Date
	}
	if m.MessageId != "" {
		metadata["message_id"] = m.MessageId
	}
	if m.ThreadId != "" {
		metadata["thread_id"] = m.ThreadId
	}
	if m.InReplyTo != "" {
		metadata["in_reply_to"] = m.InReplyTo
	}
	if m.References != "" {
		metadata["references"] = m.References
	}
	if m.Cc != "" {
		metadata["cc"] = m.Cc
	}
	if m.Bcc != "" {
		metadata["bcc"] = m.Bcc
	}
	if m.Attachments != "" {
		metadata["attachments"] = m.Attachments
	}
	if m.MimeType != "" {
		metadata["mime_type"] = m.MimeType
	}
	if m.ContentType != "" {
		metadata["content_type"] = m.ContentType
	}
	if m.Content != "" {
		metadata["content"] = m.Content
	}
	if m.Html != "" {
		metadata["html"] = m.Html
	}
	if m.Text != "" {
		metadata["text"] = m.Text
	}
	return metadata
}

func (m *FinancialTransaction) SchemaKind() string {
	return "FinancialTransaction"
}

func (m *FinancialTransaction) SchemaMetadata() map[string]interface{} {
	metadata := make(map[string]interface{})

	if m.Id != "" {
		metadata["id"] = m.Id
	}
	if m.Date != "" {
		metadata["date"] = m.Date
	}
	if m.Description != "" {
		metadata["description"] = m.Description
	}
	if m.Amount != "" {
		metadata["amount"] = m.Amount
	}
	if m.Currency != "" {
		metadata["currency"] = m.Currency
	}
	if m.Category != "" {
		metadata["category"] = m.Category
	}
	if m.Account != "" {
		metadata["account"] = m.Account
	}
	if m.Subcategory != "" {
		metadata["subcategory"] = m.Subcategory
	}
	if m.Notes != "" {
		metadata["notes"] = m.Notes
	}
	if m.Type != "" {
		metadata["type"] = m.Type
	}
	return metadata
}
