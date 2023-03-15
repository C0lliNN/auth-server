package generator

import (
	"github.com/golang-jwt/jwt"
)

type HMACSecret []byte

type JWTGenerator struct {
	secret HMACSecret
}

func NewJWTGenerator(secret HMACSecret) *JWTGenerator {
	return &JWTGenerator{secret: secret}
}

func (w *JWTGenerator) Generate() (string, error) {
	return jwt.New(jwt.SigningMethodHS256).SignedString([]byte(w.secret))
}
