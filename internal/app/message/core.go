package message

import (
	"einfach-msg/contract"
)

type (
	Message struct {
		collection contract.Collection
	}

	ResponseStatus struct {
		Status string `json:"status"`
	}
)

func New(collection contract.Collection) *Message {
	return &Message{
		collection: collection,
	}
}
