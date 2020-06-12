package utils

import (
	"math/rand"
	"time"
)

func generateRandomString(letters []byte, size int) string {
	result := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func GenerateRandomNumericText(size int) string {
	var letters = []byte("1234567890")
	return generateRandomString(letters, size)
}
