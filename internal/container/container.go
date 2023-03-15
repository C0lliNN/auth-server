package container

import (
	"C0lliNN/auth-server/internal/auth"
	"C0lliNN/auth-server/internal/config"
	"C0lliNN/auth-server/internal/generator"
	"C0lliNN/auth-server/internal/hash"
	"C0lliNN/auth-server/internal/persistence"
	"C0lliNN/auth-server/internal/server"
)

func CreateAuth(c config.Config) auth.Auth {
	db := config.NewMongoDatabase(c.Database.URI, c.Database.Name)
	idGenerator := generator.NewUUIDGenerator()
	tokenGenerator := generator.NewJWTGenerator(generator.HMACSecret(c.Token.Secret))
	hasher := hash.NewSHA256Hasher()
	clientRepo := persistence.NewClientRepository(db)
	tokenRepo := persistence.NewTokenRepository(db)
	return auth.NewAuth(auth.Config{
		ClientRepository: clientRepo,
		TokenRepository:  tokenRepo,
		TokenGenerator:   tokenGenerator,
		IDGenerator:      idGenerator,
		Hasher:           hasher,
	})
}

func CreateServer(c config.Config) server.Server {
	return server.NewServer(CreateAuth(c), c.Server.Address)
}
