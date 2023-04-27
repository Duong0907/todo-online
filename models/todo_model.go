package models

import (
	"context"
	"todo-online-api/pkgs/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title  string             `json:"title"`
	Status string             `json:"status"`
	UserID primitive.ObjectID `json:"userid,omitempty"`
}

func (t *Todo) SetUserID(userid string) error {
	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return err 
	}

	t.UserID = userID 
	return nil 
}

func (t *Todo) GetAllByUserID() ([]Todo, error) {
	result := make([]Todo, 0)

	filter := bson.M{"userid": t.UserID}
	collection := database.GetCollection("todos")
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return result, err 
	}

	err = cursor.All(context.Background(), &result)
	return result, err 
}

func (t *Todo) Create() error {
	t.ID = primitive.NewObjectID()
	collection := database.GetCollection("todos")
	_, err := collection.InsertOne(context.Background(), *t)
	return err
}

func (t *Todo) Delete(id string) error {
	collection := database.GetCollection("todos")
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": ID}
	_, err = collection.DeleteOne(context.Background(), filter)
	return err
}


func (t *Todo) ToggleCheck(id string) error {
	collection := database.GetCollection("todos")
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err 
	}
	// Find the todo having the ID
	filter := bson.M{"_id": ID}
	err = collection.FindOne(context.Background(), filter).Decode(t)
	if err != nil {
		return err 
	}

	// Determine the new status
	var newStatus string
	if t.Status == "Todo" {
		newStatus = "Done"
	} else if t.Status == "Done" {
		newStatus = "Todo"
	}
	update := bson.M{
		"$set": bson.M{
			"status": newStatus,
		},
	}
	// Update the todo having ID 
	_, err = collection.UpdateByID(context.Background(), ID, update)
	return err
}