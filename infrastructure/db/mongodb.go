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
	client  *mongo.Client
}

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

	fmt.Println(uri)

	clientOpts := options.Client().
		ApplyURI(uri)
		//SetDirect(true).
		//SetHosts([]string{"mongodb-primary", "mongodb-secondary", "mongodb-arbiter"}).
		//SetReplicaSet("replicaset")
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	//session, err := client.StartSession()
	//if err != nil {
	//	panic(err)
	//}

	return &MongoHandler{
		db: client.Database(os.Getenv("MONGODB_DATABASE")),
		//session: session,
		client: client,
	}, nil
}

func (m *MongoHandler) Client() *mongo.Client {
	return m.client
}

func (m *MongoHandler) Session() mongo.Session {
	return m.session
}

func (m *MongoHandler) Db() *mongo.Database {
	return m.db
}
