package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrExpiredToken  = errors.New("expired token")
	ErrTokenNotFound = errors.New("token not found")
)

type NotaClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateJWT(userID uuid.UUID) (string, error) {
	claims := NotaClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Second)), // TODO: change to 12 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "nota",
			Subject:   userID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		return "", errors.New("unable to find jwt secret key")
	}

	str, err := token.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return str, nil
}

func ValidateJWT(tokenStr string) (*NotaClaims, error) {
	key, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		return nil, errors.New("unable to find jwt secret key")
	}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&NotaClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}

			return []byte(key), nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return token.Claims.(*NotaClaims), ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*NotaClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
