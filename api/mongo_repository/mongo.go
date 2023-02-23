package mongo_repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	fmt.Println("Connect to MongoDB")
	return client.Database("test"), nil
}

func GetCollection(client *mongo.Database, collectionName string) *mongo.Collection {
	collection := client.Collection(collectionName)
	return collection
}
