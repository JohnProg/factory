package data

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func GenToken() string {
	sum := sha256.Sum256(GenerateRandomBytes(32))
	return base64.URLEncoding.EncodeToString(sum[:])
}

func EncryptPassword(pass string) string {
	sum := sha256.Sum256([]byte(pass))
	return base64.URLEncoding.EncodeToString(sum[:])
}
