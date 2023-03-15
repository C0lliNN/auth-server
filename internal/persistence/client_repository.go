package persistence

import (
	"C0lliNN/auth-server/internal/auth"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientRepository struct {
	db *mongo.Database
}

func NewClientRepository(db *mongo.Database) ClientRepository {
	return ClientRepository{db: db}
}

func (r ClientRepository) Save(ctx context.Context, client auth.Client) error {
	_, err := r.db.Collection("clients").ReplaceOne(ctx, bson.M{"_id": client.ID}, client, options.Replace().SetUpsert(true))
	return err
}

func (r ClientRepository) FindByID(ctx context.Context, id string) (auth.Client, error) {
	var client auth.Client
	err := r.db.Collection("clients").FindOne(ctx, bson.M{"_id": id}).Decode(&client)
	return client, err
}
