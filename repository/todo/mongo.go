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
	Database   *mongo.Database
    Collection *mongo.Collection
}

// NewMongoRepo initializes a new instance of MongoRepo.
func NewMongoRepo(client *mongo.Client) *MongoRepo {
	return &MongoRepo{
		Client: client,
		Database: client.Database("ChiTodos"),
		Collection: client.Database("ChiTodos").Collection("Todos"),
	}
}


func (m *MongoRepo) Insert(ctx context.Context, todo model.Todo) error {

	data, err := bson.Marshal(todo)
	if err != nil {
		return fmt.Errorf("failed to encode the todo into BSON %w", err)
	}
	_, err = m.Collection.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to insert the todo %w", err)
	}
	return nil
}
func (m *MongoRepo) FindAll(ctx context.Context)([]*model.Todo, error){
	filter := bson.D{}
	
	cursor, err := m.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find todos: %w", err)
	}
	defer cursor.Close(ctx)

	var todos []*model.Todo
	for cursor.Next(ctx){
		var todo model.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, fmt.Errorf("failed to decode todo : %w", err)
		}
		todos = append(todos, &todo)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error : %w", err)
	}
	
	return todos, nil

}

func (m *MongoRepo) FindByID(ctx context.Context, id uint64) (*model.Todo, error) {
    filter := bson.M{"todoid": id}

    var todo model.Todo
    err := m.Collection.FindOne(ctx, filter).Decode(&todo)
    if err != nil {
        return nil, fmt.Errorf("failed to find todo by ID: %w", err)
    }

    return &todo, nil
}
func (m *MongoRepo) Delete(ctx context.Context, id uint64) error {
    filter := bson.M{"todoid": id}

    _, err := m.Collection.DeleteOne(ctx, filter)
    if err != nil {
        return fmt.Errorf("failed to delete todo: %w", err)
    }

    return nil
}
func (m *MongoRepo) Update(ctx context.Context, id uint64, update bson.M) error {
    filter := bson.M{"todoid": id}

    _, err := m.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
    if err != nil {
        return fmt.Errorf("failed to update todo: %w", err)
    }

    return nil
}