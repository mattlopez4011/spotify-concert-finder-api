package models

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

const numRandomBytes = 32

func GenerateRandomString() (string, error) {
	b, err := GenerateRandomBytes(numRandomBytes)
	if err != nil {
		return "", err
	}

	return EncodeBase64WithoutPadding(b), nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func EncodeBase64WithoutPadding(input []byte) string {
	encoded := base64.URLEncoding.EncodeToString(input)
	parts := strings.Split(encoded, "=")
	encoded = parts[0]

	return encoded
}
