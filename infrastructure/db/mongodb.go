package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHandler struct {
	db      *mongo.Database
	session mongo.Session
}

func NewMongoHandler() (*MongoHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	uri := fmt.Sprintf(
		"%s://@%s",
		os.Getenv("MONGODB_HOST"),
		os.Getenv("MONGODB_HOST"),
	)

	clientOpts := options.Client().ApplyURI(uri).SetDirect(true)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}

	return &MongoHandler{
		db:      client.Database(os.Getenv("MONGODB_DATABASE")),
		session: session,
	}, nil
}

func (m *MongoHandler) Session() mongo.Session {
	return m.session
}

func (m *MongoHandler) Db() *mongo.Database {
	return m.db
}
