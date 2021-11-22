package hash

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GenerateToken(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text+time.Now().String()), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash), nil
}
