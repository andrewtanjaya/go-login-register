package utils

import (
	"fmt"
	"go-login-register/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := JWTClaims{
		user.ID,
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(jwtSecretKey)

	return jwt, err
}

func ValidateToken(stringToken string) (any, error) {
	token, err := jwt.ParseWithClaims(stringToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, success := token.Claims.(*JWTClaims)
	if !success || !token.Valid {
		return nil, fmt.Errorf("unauthorized: invalid token")
	}

	return claims, nil
}
