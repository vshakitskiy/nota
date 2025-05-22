package bcrypt

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"nota.shared/telemetry"
)

func Hash(ctx context.Context, password string) (string, error) {
	ctx, span := telemetry.StartSpan(ctx, "Bcrypt.Hash")
	defer span.End()

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func Compare(ctx context.Context, password, hash string) bool {
	ctx, span := telemetry.StartSpan(ctx, "Bcrypt.Compare")
	defer span.End()

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}
