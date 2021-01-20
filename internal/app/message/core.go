package message

import (
	"go-simple-svc/contract"
)

type (
	// Message is a compulsory dependencies or configuration for message service
	Message struct {
		collection contract.Collector
	}

	// ResponseStatus used as http response model in create action
	ResponseStatus struct {
		Status string `json:"status"`
	}
)

// New will instantiate Message
func New(collection contract.Collector) *Message {
	return &Message{
		collection: collection,
	}
}
