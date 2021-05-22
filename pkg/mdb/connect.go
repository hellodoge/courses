package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	connectionTimeout = 5 * time.Second
)

func Connect(uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, client.Ping(ctx, readpref.Primary())
}
