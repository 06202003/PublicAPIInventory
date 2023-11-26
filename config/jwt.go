package config

import (
	"crypto/rand"
	"log"
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY []byte

// GenerateRandomKey generates a random key with the specified length
func GenerateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return key, nil
}

// GetJWTKey returns the JWT key
func GetJWTKey() []byte {
	// Load the key from a secure location or configuration
	// For simplicity, we use a static key in development/testing
	if len(JWT_KEY) == 0 {
		log.Fatal("Kunci JWT belum diatur. Silakan atur Secure Key.")
	}
	return JWT_KEY
}

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
