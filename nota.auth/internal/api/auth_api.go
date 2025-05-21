package api

import (
	"context"
	"log"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nota.auth/internal/repository"
	"nota.auth/internal/service"
	pb "nota.auth/pkg/pb/v1"
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
	panic("not implemented")
}

func (h *AuthServiceHandler) Login(
	ctx context.Context,
	req *pb.LoginRequest,
) (*pb.LoginResponse, error) {
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
		switch err {
		case service.ErrIncorrectPassword:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case repository.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			log.Println(err)
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

func (h *AuthServiceHandler) RefreshAccessToken(
	ctx context.Context,
	req *pb.RefreshTokenRequest,
) (*pb.RefreshTokenResponse, error) {
	panic("not implemented")
}

func (h *AuthServiceHandler) Register(
	ctx context.Context,
	req *pb.RegisterRequest,
) (*pb.RegisterResponse, error) {
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
		switch err {
		case service.ErrInvalidPassword:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case repository.ErrUserAlreadyExists:
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
