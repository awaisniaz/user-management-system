package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret     = []byte("AccessSecret")
	refreshSecret = []byte("yourrefreshSecret")
)

func GenerateAccessToken(userId, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtSecret)

}

func GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}
