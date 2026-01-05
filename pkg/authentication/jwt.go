package authentication

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(secret, username, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenString string) (string, string, error) {
	tokenString = strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer"))
	if tokenString == "" {
		return "", "", errors.New("Token Invalid")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), claims["email"].(string), nil
	}
	return "", "", err
}
