package api

import (
	"context"

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
	panic("not implemented")
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
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"username, email and password are required",
		)
	}

	id, err := h.service.User.Create(ctx, req.Username, req.Email, req.Password)
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
