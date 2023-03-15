package hash

import "crypto/sha256"

type SHA256Hasher struct{}

func NewSHA256Hasher() SHA256Hasher {
	return SHA256Hasher{}
}

func (SHA256Hasher) Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))

	hashed := h.Sum(nil)
	return string(hashed)
}
