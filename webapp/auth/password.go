package auth

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(password string) string {
	h := md5.New()
	hashedPassword := h.Sum([]byte(password))

	return hex.EncodeToString(hashedPassword)
}
