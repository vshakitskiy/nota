package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"nota.auth/internal/model"
	"nota.auth/internal/repository"
	"nota.auth/pkg/bcrypt"
	"nota.auth/pkg/crypto"
	"nota.auth/pkg/jwt"
	"nota.shared/telemetry"
)

var (
	ErrInvalidPassword   = errors.New("password is invalid")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrSessionExpired    = errors.New("session expired")
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type AuthService interface {
	Register(ctx context.Context, username, email, password string) (*uuid.UUID, error)
	Login(ctx context.Context, email, password string) (*TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	Logout(ctx context.Context, accessToken string) error
	GetUser(ctx context.Context, accessToken string) (*model.User, error)
}

type AuthServiceImpl struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthServiceImpl {
	return &AuthServiceImpl{
		repo: repo,
	}
}

func (s *AuthServiceImpl) Register(
	ctx context.Context,
	username, email, password string,
) (*uuid.UUID, error) {
	// TODO: ensure email is not already in use

	ctx, span := telemetry.StartSpan(ctx, "AuthService.Register")
	defer span.End()

	user := &model.User{
		Username: username,
		Email:    email,
	}

	passwordHash, err := bcrypt.Hash(ctx, password)
	if err != nil {
		telemetry.RecordError(span, err)
		return nil, ErrInvalidPassword
	}
	user.Password = passwordHash

	userID, err := s.repo.User.Create(ctx, user)
	if err != nil {
		telemetry.RecordError(span, err)
		return nil, err
	}

	return userID, nil
}

func (s *AuthServiceImpl) Login(
	ctx context.Context,
	email, password string,
) (*TokenPair, error) {
	ctx, span := telemetry.StartSpan(ctx, "AuthService.Login")
	defer span.End()

	user, err := s.repo.User.GetByEmail(ctx, email)
	if err != nil {
		telemetry.RecordError(span, err)
		return nil, err
	}

	ok := bcrypt.Compare(ctx, password, user.Password)
	if !ok {
		telemetry.RecordError(span, ErrIncorrectPassword)
		return nil, ErrIncorrectPassword
	}

	accessToken, err := jwt.CreateJWT(user.ID)
	if err != nil {
		telemetry.RecordError(span, err)
		return nil, err
	}
	refreshToken := crypto.GenerateRandomBase64(64)

	session := &model.Session{
		RefreshToken: refreshToken,
		UserID:       user.ID,
		ExpiresAt:    time.Now().Add(12 * 24 * time.Hour),
	}

	if err := s.repo.Session.Create(ctx, session); err != nil {
		telemetry.RecordError(span, err)
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceImpl) RefreshToken(
	ctx context.Context,
	refreshToken string,
) (string, error) {
	ctx, span := telemetry.StartSpan(ctx, "AuthService.RefreshToken")
	defer span.End()

	session, err := s.repo.Session.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		telemetry.RecordError(span, err)
		return "", err
	}

	if session.ExpiresAt.Before(time.Now()) {
		err := s.repo.Session.DeleteExpiredByUserID(ctx, session.UserID)
		if err != nil {
			telemetry.RecordError(span, err)
			return "", err
		}

		return "", ErrSessionExpired
	}

	accessToken, err := jwt.CreateJWT(session.UserID)
	if err != nil {
		telemetry.RecordError(span, err)
		return "", err
	}

	return accessToken, nil
}

func (s *AuthServiceImpl) Logout(
	ctx context.Context,
	accessToken string,
) error {
	ctx, span := telemetry.StartSpan(ctx, "AuthService.Logout")
	defer span.End()

	claims, err := jwt.ValidateJWT(accessToken)
	if err != nil {
		telemetry.RecordError(span, err)
		return err
	}

	return s.repo.Session.DeleteByUserID(ctx, claims.UserID)
}

func (s *AuthServiceImpl) GetUser(
	ctx context.Context,
	accessToken string,
) (*model.User, error) {
	ctx, span := telemetry.StartSpan(ctx, "AuthService.GetUser")
	defer span.End()

	claims, err := jwt.ValidateJWT(accessToken)
	if err != nil {
		telemetry.RecordError(span, err)
		return nil, err
	}

	return s.repo.User.GetById(ctx, claims.UserID)
}
