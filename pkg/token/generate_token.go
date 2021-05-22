package token

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/hellodoge/courses-tg-bot/courses"
	"io"
)

func GenerateToken() (string, error) {
	tokenBytesLength := courses.TokenLength / 2
	token := make([]byte, tokenBytesLength)
	_, err := io.ReadFull(rand.Reader, token)
	return hex.EncodeToString(token), err
}
