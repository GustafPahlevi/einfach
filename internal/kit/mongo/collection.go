package mongo

import (
	"context"
	"einfach-msg/constant"
	"einfach-msg/internal/kit/datetime"
	"einfach-msg/model"
	"time"

	"github.com/jonboulle/clockwork"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	collection *mongo.Collection
	dbTimeout  time.Duration
}

func New(collection *mongo.Collection, dbTimeout time.Duration) *Collection {
	return &Collection{
		collection: collection,
		dbTimeout:  dbTimeout,
	}
}

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
