package mongo

import (
	"context"
	"fmt"
	"go-simple-svc/constant"
	"go-simple-svc/internal/kit/datetime"
	"go-simple-svc/model"
	"time"

	"github.com/jonboulle/clockwork"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection is a compulsory dependencies or configuration for mongodb collection
type Collection struct {
	collection *mongo.Collection
	dbTimeout  time.Duration
}

// New will instantiate mongodb collection
func New(collection *mongo.Collection, dbTimeout time.Duration) *Collection {
	return &Collection{
		collection: collection,
		dbTimeout:  dbTimeout,
	}
}

// Insert document to mongodb
func (c *Collection) Insert(request model.Model) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.dbTimeout*time.Second)
	defer cancel()

	clock := clockwork.NewRealClock()
	_, err := c.collection.InsertOne(ctx, bson.M{
		"sender_id":   request.SenderID,
		"receiver_id": request.ReceiverID,
		"subject":     request.Subject,
		"message":     request.Message,
		"status":      constant.Success,
		"time":        datetime.GetCurrentDateTime(clock),
	})
	if err != nil {
		return err
	}

	return nil
}

// Get document from mongodb
func (c *Collection) Get() ([]*model.Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.dbTimeout*time.Second)
	defer cancel()
	var messages []*model.Model

	cursor, err := c.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var message model.Model
		err = cursor.Decode(&message)
		if err != nil {
			return nil, err
		}
		fmt.Println("msg_id: ", message.ID)
		fmt.Println("msg: ", message.ReceiverID)
		fmt.Println("msg: ", message.SenderID)
		fmt.Println("msg: ", message.Message)
		fmt.Println("msg: ", message.Status)
		messages = append(messages, &message)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	err = cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
