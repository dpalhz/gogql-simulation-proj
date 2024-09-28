package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateUUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(fmt.Errorf("failed to generate UUID: %v", err))
	}
	return hex.EncodeToString(uuid)
}
