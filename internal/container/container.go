package container

import (
	"C0lliNN/auth-server/internal/auth"
	"C0lliNN/auth-server/internal/generator"
	"C0lliNN/auth-server/internal/hash"
	"C0lliNN/auth-server/internal/persistence"
)

func CreateAuth() auth.Auth {
	idGenerator := generator.NewUUIDGenerator()
	hasher := hash.NewSHA256Hasher()
	clientRepo := persistence.NewClientRepository()
	return auth.NewAuth(auth.Config{
		ClientRepository: clientRepo,
		TokenRepository:  nil,
		IDGenerator:      idGenerator,
		Hasher:           hasher,
	})
}
