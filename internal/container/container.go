package container

import (
	"C0lliNN/auth-server/internal/auth"
	"C0lliNN/auth-server/internal/generator"
	"C0lliNN/auth-server/internal/hash"
	"C0lliNN/auth-server/internal/persistence"
)

func CreateAuth() auth.Auth {
	idGenerator := generator.NewUUIDGenerator()
	tokenGenerator := generator.NewJWTGenerator([]byte("my secret"))
	hasher := hash.NewSHA256Hasher()
	clientRepo := persistence.NewClientRepository()
	tokenRepo := persistence.NewTokenRepository()
	return auth.NewAuth(auth.Config{
		ClientRepository: clientRepo,
		TokenRepository:  tokenRepo,
		TokenGenerator:   tokenGenerator,
		IDGenerator:      idGenerator,
		Hasher:           hasher,
	})
}
