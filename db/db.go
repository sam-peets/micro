package db

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DbConnection struct {
	Client mongo.Client
}

func (conn *DbConnection) Context() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return ctx, cancel
}

func (conn *DbConnection) Close() {
	log.Println("closing db connection")
	ctx, cancel := conn.Context()
	defer cancel()
	err := conn.Client.Disconnect(ctx)
	if err != nil {
		log.Panic(err)
	}
}

func (conn *DbConnection) GetCollection(name string) (*mongo.Collection, error) {
	coll := conn.Client.Database("micro").Collection(name)
	if coll == nil {
		return nil, errors.New("couldn't get collection")
	}
	return coll, nil
}

var client *DbConnection = nil

func GetClient() *DbConnection {
	if client != nil {
		return client
	}
	log.Println("mongo connection is nil, connecting")
	mongoClient, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Panicln(err)
	}
	client = &DbConnection{
		Client: *mongoClient,
	}
	return client
}
