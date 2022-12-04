package models

import "time"

type CreateMessageRequest struct {
    Sender int64 `example:"1" json"sender" validate:"required"` // Sending user
    Recipient int64 `example:"2" json"recipient" validate:"required"` // Recieving user

    MessageData
}

type GeMessagesResponse struct {
	Items []*MessageItem // Found items
	*Page
}

type MessageData struct {
    Text string `example:"A message that need to be sent" json:"text" validate:"required"` // Message text
}

type MessageItem struct {
    ID string `json:"id" validate:"required"` // Message id
    Sent time.Time `json:"sent" validate:"required"` // Message creation time
    Edited time.Time `json:"edited" validate:"required"` // Message last edit time
    IsReceived bool `json:"is_received" validate:"required"` // Whether message was received

    Sender int64 `example:"1" json"sender" validate:"required"` // Sending user
    Recipient int64 `example:"2" json"recipient" validate:"required"` // Recieving user
    Text string `json:"text" validate:"required"` // Message text
    OriginalID* string `json"original_id"` // Original message id
}
