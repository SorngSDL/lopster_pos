package db

import (
	"context"
	"log"
	"time"

	configs "com.lopster-pos/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(configs.MongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Connect mongoDB failed: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Ping mongoDB failed: %v", err)
	}

	Client = client
	log.Println("Connect mongoDB successful!")
}

func GetCollection(name string) *mongo.Collection {
	return Client.Database(configs.DBName).Collection(name)
}
