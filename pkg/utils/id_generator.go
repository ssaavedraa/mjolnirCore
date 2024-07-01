package utils

import (
	"math/rand"
	"time"
)

func GenerateId() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.NewSource(time.Now().UnixNano())

	b := make([]byte, 16)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
