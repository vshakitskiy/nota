package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type NotaClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateJWT(userID uuid.UUID) (string, error) {
	claims := NotaClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
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
	claims := new(NotaClaims)

	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return t, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
