package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client , err := mongo.Connect(ctx, options.Client().ApplyURI(LoadEnv("MONGO_URI")))
	
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	}

	fmt.Print("Connected to MongoDB\n")
	return client

}

var Client *mongo.Client = MongoConnect()

func GetCOllection (client *mongo.Client, collectionName string) *mongo.Collection{
	return client.Database(LoadEnv("DB_NAME")).Collection(collectionName)
}