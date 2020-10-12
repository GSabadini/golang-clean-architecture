package db

import (
	"context"
	"fmt"
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

	uri := fmt.Sprintf(
		"%s://root:password123@%s,%s,%s/?replicaSet=%s",
		os.Getenv("MONGODB_HOST"),
		"mongodb-primary",
		"mongodb-secondary",
		"mongodb-arbiter",
		"replicaset",
	)

	clientOpts := options.Client().ApplyURI(uri)
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
