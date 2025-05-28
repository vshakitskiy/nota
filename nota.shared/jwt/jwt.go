package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"nota.shared/env"
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

func CreateJWT(userID uuid.UUID, exp time.Duration) (string, error) {
	claims := NotaClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "nota",
			Subject:   userID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := env.GetJwtSecret()

	str, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return str, nil
}

func ValidateJWT(tokenStr string) (*NotaClaims, error) {
	secret := env.GetJwtSecret()

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&NotaClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}

			return []byte(secret), nil
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
