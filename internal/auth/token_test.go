package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToken_ExpiresIn(t *testing.T) {
	token := Token{ExpirationTime: time.Now().Add(time.Second * 3600).Unix()}
	assert.Equal(t, int64(3600), token.ExpiresIn())
}
