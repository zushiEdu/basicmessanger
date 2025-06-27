package smallFunctions

import (
	"crypto/rand"
	"encoding/hex"
	"main/config"
)

func GenerateToken() string {
	b := make([]byte, config.TokenLength)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
