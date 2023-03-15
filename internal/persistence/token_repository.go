package persistence

import (
	"C0lliNN/auth-server/internal/auth"
	"context"
)

type TokenRepository struct{}

func NewTokenRepository() TokenRepository {
	return TokenRepository{}
}

func (r TokenRepository) Save(ctx context.Context, token auth.Token) error {
	return nil
}
