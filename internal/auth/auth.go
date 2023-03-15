package auth

import (
	"context"
	"fmt"
	"time"
)

type ClientRepository interface {
	Save(ctx context.Context, c Client) error
	FindByID(ctx context.Context, id string) (Client, error)
}

type TokenRepository interface {
	Save(ctx context.Context, t Token) error
}

type IDGenerator interface {
	NewID() string
}

type TokenGenerator interface {
	Generate(map[string]interface{}) (string, error)
}

type Hasher interface {
	Hash(string) string
}

type Config struct {
	ClientRepository ClientRepository
	TokenRepository  TokenRepository
	IDGenerator      IDGenerator
	TokenGenerator   TokenGenerator
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

type ObtainTokenRequest struct {
	GrantType    Grant  `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type ClientResponse struct {
	ID          string
	Type        string
	RedirectURI *string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
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
		hashedPassword := a.Hasher.Hash(req.Secret)
		client.Secret = &hashedPassword
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

func (a Auth) ObtainToken(ctx context.Context, req ObtainTokenRequest) (TokenResponse, error) {
	if err := a.validateObtainTokenRequest(req); err != nil {
		return TokenResponse{}, err
	}

	client, err := a.ClientRepository.FindByID(ctx, req.ClientID)
	if err != nil {
		return TokenResponse{}, err
	}

	if *client.Secret != a.Hasher.Hash(req.ClientSecret) {
		return TokenResponse{}, fmt.Errorf("invalid secret")
	}

	expirationTime := time.Now().Add(time.Hour).Unix()

	accessToken, err := a.TokenGenerator.Generate(map[string]interface{}{
		"client_id": client.ID,
		"exp":       expirationTime,
	})
	if err != nil {
		return TokenResponse{}, err
	}

	token := Token{
		ID:             a.IDGenerator.NewID(),
		Type:           "Bearer",
		AccessToken:    accessToken,
		CreatedAt:      time.Now().Unix(),
		ExpirationTime: expirationTime,
	}

	if err = a.TokenRepository.Save(ctx, token); err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		AccessToken:  token.AccessToken,
		TokenType:    token.Type,
		ExpiresIn:    token.ExpiresIn(),
		RefreshToken: token.RefreshToken,
	}, nil
}

func (a Auth) validateObtainTokenRequest(req ObtainTokenRequest) error {
	if req.GrantType == "" {
		return fmt.Errorf("the grant_type is required")
	}

	if req.ClientID == "" {
		return fmt.Errorf("the client_id is required")
	}

	if req.GrantType == ClientCredentials && req.ClientSecret == "" {
		return fmt.Errorf("the client_secret must be provided when grant_type is client_credentials")
	}

	return nil
}
