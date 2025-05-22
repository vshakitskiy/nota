package api

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nota.auth/internal/repository"
	"nota.auth/internal/service"
	"nota.auth/pkg/jwt"
	pb "nota.auth/pkg/pb/v1"
	"nota.shared/telemetry"
)

type AuthServiceHandler struct {
	pb.UnimplementedAuthServiceServer
	service *service.Service
}

func NewAuthServiceHandler(service *service.Service) *AuthServiceHandler {
	return &AuthServiceHandler{
		service: service,
	}
}

func (h *AuthServiceHandler) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {
	if req.AccessToken == "" {
		return nil, status.Error(codes.InvalidArgument, "access token is required")
	}

	user, err := h.service.Auth.GetUser(ctx, req.AccessToken)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrInvalidToken):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, jwt.ErrExpiredToken):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, repository.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(
				codes.Internal,
				"something went wrong, try again later",
			)
		}
	}

	return &pb.GetUserResponse{
		UserId:   user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (h *AuthServiceHandler) Login(
	ctx context.Context,
	req *pb.LoginRequest,
) (*pb.LoginResponse, error) {
	ctx, span := telemetry.StartSpan(ctx, "AuthHandler.Login")
	defer span.End()

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if req.Email == "" || !emailRegex.MatchString(req.Email) || len(req.Email) > 254 {
		return nil, status.Error(
			codes.InvalidArgument,
			"email is required and must be valid",
		)
	}

	if req.Password == "" || len(req.Password) < 8 || len(req.Password) > 20 {
		return nil, status.Error(
			codes.InvalidArgument,
			"password is required and must be between 8 and 20 characters long",
		)
	}

	tokenPair, err := h.service.Auth.Login(ctx, req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrIncorrectPassword):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, repository.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(
				codes.Internal,
				"something went wrong, try again later",
			)
		}
	}

	return &pb.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (h *AuthServiceHandler) RefreshToken(
	ctx context.Context,
	req *pb.RefreshTokenRequest,
) (*pb.RefreshTokenResponse, error) {
	ctx, span := telemetry.StartSpan(ctx, "AuthHandler.RefreshToken")
	defer span.End()

	if req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh token is required")
	}

	accessToken, err := h.service.Auth.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrSessionExpired):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, repository.ErrSessionNotFound):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Error(
				codes.Internal,
				"something went wrong, try again later",
			)
		}
	}

	return &pb.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (h *AuthServiceHandler) Register(
	ctx context.Context,
	req *pb.RegisterRequest,
) (*pb.RegisterResponse, error) {
	ctx, span := telemetry.StartSpan(ctx, "AuthHandler.Register")
	defer span.End()

	if req.Username == "" || len(req.Username) < 3 || len(req.Username) > 20 {
		return nil, status.Error(
			codes.InvalidArgument,
			"username is required and must be between 3 and 20 characters long",
		)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if req.Email == "" || !emailRegex.MatchString(req.Email) || len(req.Email) > 254 {
		return nil, status.Error(
			codes.InvalidArgument,
			"email is required and must be valid",
		)
	}

	if req.Password == "" || len(req.Password) < 8 || len(req.Password) > 20 {
		return nil, status.Error(
			codes.InvalidArgument,
			"password is required and must be between 8 and 20 characters long",
		)
	}

	id, err := h.service.Auth.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidPassword):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, repository.ErrUserAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(
				codes.Internal,
				"something went wrong, try again later",
			)
		}
	}

	return &pb.RegisterResponse{
		UserId: id.String(),
	}, nil
}

func (h *AuthServiceHandler) Logout(
	ctx context.Context,
	req *pb.LogoutRequest,
) (*pb.LogoutResponse, error) {
	ctx, span := telemetry.StartSpan(ctx, "AuthHandler.Logout")
	defer span.End()

	if req.AccessToken == "" {
		return nil, status.Error(codes.InvalidArgument, "access token is required")
	}

	err := h.service.Auth.Logout(ctx, req.AccessToken)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrExpiredToken):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			fmt.Println(err)
			return nil, status.Error(
				codes.Internal,
				"something went wrong, try again later",
			)
		}
	}

	return &pb.LogoutResponse{}, nil
}
