package auth

import "time"

type Token struct {
	ID             string
	Type           string
	AccessToken    string
	RefreshToken   string
	State          string
	Scope          string
	CreatedAt      int64
	ExpirationTime int64
}

func (t Token) ExpiresIn() int64 {
	return t.ExpirationTime - time.Now().Unix()
}
