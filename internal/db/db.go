package db

import (
	"context"
	"log"

	"project/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Instance *Db
)

// Db struct
type Db struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// New creates db connection
func New(config *config.Config) *Db {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoDBUrl)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return &Db{
		Client:   client,
		Database: client.Database(config.DBName),
	}
}
