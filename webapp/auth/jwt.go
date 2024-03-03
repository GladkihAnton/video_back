package auth

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"streamer/config"
	"streamer/webapp/db/models"
	"time"
)

type JwtToken struct {
	AccessToken string `json:"access_token"`
	Exp         int64  `json:"exp"`
}

func GenerateToken(user *models.User) ([]byte, error) {
	exp := time.Now().Add(time.Hour * 72).Unix()
	payload := jwt.MapClaims{
		"user_id": user.Id,
		"exp":     exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	accessToken, jwtTokenErr := token.SignedString(config.JwtSecretKey)
	if jwtTokenErr != nil {
		return nil, jwtTokenErr
	}
	responseToken, parseTokenErr := json.Marshal(JwtToken{
		AccessToken: accessToken,
		Exp:         exp,
	})
	if parseTokenErr != nil {
		return nil, parseTokenErr
	}

	return responseToken, nil
}
