package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateUsername() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("user_%x", b)
}
