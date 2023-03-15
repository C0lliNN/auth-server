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

func (w *JWTGenerator) Generate(claims map[string]interface{}) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims)).SignedString([]byte(w.secret))
}
