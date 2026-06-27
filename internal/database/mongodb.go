package database

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func Connect() (*mongo.Client, error) {
	client, err := mongo.Connect(
		options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_URL")),
	)

	if err != nil {
		return nil, err
	}

	DB = client.Database("productdb")
	fmt.Println("Connected to ", DB.Name())

	return client, nil
}