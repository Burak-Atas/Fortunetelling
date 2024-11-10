package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client = SetDB()

func SetDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(""))
	if err != nil {
		log.Fatal(err)
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client.Ping(ctx, readpref.Primary())

	return client
}

func UserCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	return client.Database("FALCI_BACI").Collection(collectionName)
}
