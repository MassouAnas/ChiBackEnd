// mongo.go

package todo

import (
	"context"
	"fmt"

	"github.com/MassouAnas/ChiBackEnd/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	Client *mongo.Client
}

// NewMongoRepo initializes a new instance of MongoRepo.
func NewMongoRepo(client *mongo.Client) *MongoRepo {
	return &MongoRepo{
		Client: client,
	}
}

func (m *MongoRepo) Insert(ctx context.Context, todo model.Todo) error {
	database := m.Client.Database("ChiTodos")
	collection := database.Collection("Todos")
	data, err := bson.Marshal(todo)
	if err != nil {
		return fmt.Errorf("failed to encode the todo into BSON %w", err)
	}

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to insert the todo %w", err)
	}
	return nil
}
