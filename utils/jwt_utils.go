package utils

import (
	"fmt"
	"time"

	"example.com/startup/internal/dtos"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("12345678901234567890123456789012")

func GenerateToken(claims dtos.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

func GenerateEmptyToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString(JWTSecret)
}

func ExtractClaims(signedToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid claims")
}

func ValidateToken(signedToken string) bool {
	_, err := ExtractClaims(signedToken)
	return err == nil
}

func GenerateRegisteredClaims(issuer string, expiryDuration time.Duration) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiryDuration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    issuer,
	}
}
