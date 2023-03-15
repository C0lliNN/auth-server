package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(uri string, dbname string) *mongo.Database {
	ctx := context.Background()
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err = mongoClient.Connect(ctx); err != nil {
		panic(err)
	}

	if err = mongoClient.Ping(ctx, nil); err != nil {
		panic(err)
	}

	return mongoClient.Database(dbname)
}
