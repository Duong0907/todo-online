package database 

import (
	"sync"
	"os"	
	"context"
	"time"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once 
var client *mongo.Client
var database *mongo.Database

func GetCollection(collectionName string) *mongo.Collection {
	once.Do(func() {
		uri := os.Getenv("MONGODB_CONNECTION_STRING")
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()


		var err error
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	})

	databaseName := os.Getenv("MONGODB_DATABASE_NAME")
	database = client.Database(databaseName)
	collection := database.Collection(collectionName)

	return collection
}