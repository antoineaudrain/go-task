package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateSecretKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("failed to generate secret key: %v", err)
	}

	return base64.StdEncoding.EncodeToString(key)
}
