package persistence

import (
	"C0lliNN/auth-server/internal/auth"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenRepository struct {
	db *mongo.Database
}

func NewTokenRepository(db *mongo.Database) TokenRepository {
	return TokenRepository{db: db}
}

func (r TokenRepository) Save(ctx context.Context, token auth.Token) error {
	_, err := r.db.Collection("clients").ReplaceOne(ctx, bson.M{"_id": token.ID}, token, options.Replace().SetUpsert(true))
	return err
}
