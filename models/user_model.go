package models

import (
	"context"
	"log"

	"todo-online-api/pkgs/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

func (u *User) Create(username, email, password string) error {
	*u = User{
		primitive.NewObjectID(),
		username,
		email,
		password,
	}

	collection := database.GetCollection("users")
	_, err := collection.InsertOne(context.Background(), *u)
	log.Println(err)
	return err
}

func (u *User) GetOne() error {
	collection := database.GetCollection("users")
	filter := bson.M{"_id": u.ID}
	log.Println(filter)
	err := collection.FindOne(context.Background(), filter).Decode(u)
	return err
}

func (u *User) GetByUsername(username string) error {
	collection := database.GetCollection("users")
	filter := bson.M{
		"username": username,
	}
	err := collection.FindOne(context.Background(), filter).Decode(u)
	return err
}

func (u *User) Existed(username, email string) bool {
	collection := database.GetCollection("users")

	name_condition := bson.M{"username": username}
	email_condition := bson.M{"email": email}
	name_count, err := collection.CountDocuments(context.Background(), name_condition)
	if err != nil {
		return true
	}
	email_count, err := collection.CountDocuments(context.Background(), email_condition)
	if err != nil {
		return true
	}
	return name_count > 0 || email_count > 0
}
