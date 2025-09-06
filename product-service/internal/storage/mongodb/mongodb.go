package mongodb

import (
	"context"
	"fmt"
	"product-service/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Client     mongo.Client
	Collection mongo.Collection
}

func NewConnection(cfg *config.Config) (*Mongo, error) {

	uri := fmt.Sprintf("mongodb://%s:%d", cfg.Mongo.Host, cfg.Mongo.Port)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}
	mycoll := client.Database(cfg.Mongo.Database).Collection(cfg.Mongo.Collection)
	return &Mongo{
		Client:     *client,
		Collection: *mycoll,
	}, nil
}
