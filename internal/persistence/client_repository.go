package persistence

import (
	"C0lliNN/auth-server/internal/auth"
	"context"
)

type ClientRepository struct{}

func NewClientRepository() ClientRepository {
	return ClientRepository{}
}

func (r ClientRepository) Save(ctx context.Context, client auth.Client) error {
	return nil
}

func (r ClientRepository) FindByID(ctx context.Context, id string) (auth.Client, error) {
	return auth.Client{}, nil
}
