package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoHandler defines the MongoDb handler
type MongoHandler struct {
	db     *mongo.Database
	client *mongo.Client
}

// NewMongoHandler creates new MongoHandler
func NewMongoHandler() (*MongoHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return &MongoHandler{
		db:     client.Database(os.Getenv("MONGODB_DATABASE")),
		client: client,
	}, nil
}

// Client returns the client property
func (m *MongoHandler) Client() *mongo.Client {
	return m.client
}

// Db returns the db property
func (m *MongoHandler) Db() *mongo.Database {
	return m.db
}
