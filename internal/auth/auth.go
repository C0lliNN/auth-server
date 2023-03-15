package auth

import (
	"context"
	"fmt"
)

type ClientRepository interface {
	Save(ctx context.Context, c Client) error
	FindByID(ctx context.Context, id string) (Client, error)
}

type TokenRepository interface {
	Save(ctx context.Context, c Client) error
}

type IDGenerator interface {
	NewID() string
}

type Hasher interface {
	Hash(string) string
}

type Config struct {
	ClientRepository ClientRepository
	TokenRepository  TokenRepository
	IDGenerator      IDGenerator
	Hasher           Hasher
}

type Auth struct {
	Config
}

type CreateClientRequest struct {
	Type        ClientType
	Secret      string
	RedirectURI string
}

type ClientResponse struct {
	ID          string
	Type        string
	RedirectURI *string
}

func NewAuth(c Config) Auth {
	return Auth{c}
}

func (a Auth) CreateNewClient(ctx context.Context, req CreateClientRequest) (ClientResponse, error) {
	if err := a.validateCreateClientRequest(req); err != nil {
		return ClientResponse{}, err
	}

	client := Client{
		ID:   a.IDGenerator.NewID(),
		Type: req.Type,
	}

	if client.Type == Public {
		client.RedirectURI = &req.RedirectURI
	}

	if client.Type == Confidential {
		client.Secret = &req.Secret
	}

	if err := a.ClientRepository.Save(ctx, client); err != nil {
		return ClientResponse{}, err
	}

	return a.createClientResponse(client), nil
}

func (a Auth) validateCreateClientRequest(req CreateClientRequest) error {
	if req.Type != Public && req.Type != Confidential {
		return fmt.Errorf("invalid type")
	}

	if req.Type == Confidential && req.Secret == "" {
		return fmt.Errorf("confidential clients must have a secret")
	}

	if req.Type == Public && req.RedirectURI == "" {
		return fmt.Errorf("public clients must have a redirect uri")
	}

	return nil
}

func (a Auth) createClientResponse(client Client) ClientResponse {
	return ClientResponse{
		ID:          client.ID,
		Type:        client.Type.String(),
		RedirectURI: client.RedirectURI,
	}
}
