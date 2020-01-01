package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	SenderID   string             `json:"sender_id"`
	ReceiverID string             `json:"receiver_id"`
	Subject    string             `json:"subject"`
	Message    string             `json:"message"`
	Status     string             `json:"status"`
	Time       string             `json:"time"`
}
