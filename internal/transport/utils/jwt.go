package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json: "email"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("6b60e7ada6a1571fc2473150c9283235d69062f7c32891434af424075d9277a9")

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	t, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
