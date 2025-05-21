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
)

var (
	ErrInvalidPassword   = errors.New("password is invalid")
	ErrIncorrectPassword = errors.New("incorrect password")
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type AuthService interface {
	Register(ctx context.Context, username, email, password string) (*uuid.UUID, error)
	Login(ctx context.Context, email, password string) (*TokenPair, error)
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
	user := &model.User{
		Username: username,
		Email:    email,
	}

	passwordHash, err := bcrypt.Hash(password)
	if err != nil {
		return nil, ErrInvalidPassword
	}
	user.Password = passwordHash

	return s.repo.User.Create(ctx, user)
}

func (s *AuthServiceImpl) Login(
	ctx context.Context,
	email, password string,
) (*TokenPair, error) {
	user, err := s.repo.User.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	ok := bcrypt.Compare(password, user.Password)
	if !ok {
		return nil, ErrIncorrectPassword
	}

	accessToken, err := jwt.CreateJWT(user.ID)
	if err != nil {
		return nil, err
	}
	refreshToken := crypto.GenerateRandomBase64(64)

	session := &model.Session{
		RefreshToken: refreshToken,
		UserID:       user.ID,
		ExpiresAt:    time.Now().Add(12 * 24 * time.Hour),
	}

	if err := s.repo.Session.Create(ctx, session); err != nil {
		return nil, err
	}

	tokenPair := new(TokenPair)
	tokenPair.AccessToken = accessToken
	tokenPair.RefreshToken = refreshToken

	return tokenPair, nil
}
