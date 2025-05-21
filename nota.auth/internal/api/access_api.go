package api

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nota.auth/internal/service"
	"nota.auth/pkg/jwt"
	pb "nota.auth/pkg/pb/v1"
)

type AccessServiceHandler struct {
	pb.UnimplementedAccessServiceServer
	service *service.Service
}

func NewAccessServiceHandler(service *service.Service) *AccessServiceHandler {
	return &AccessServiceHandler{
		service: service,
	}
}

func (h *AccessServiceHandler) ValidateToken(
	ctx context.Context,
	req *pb.ValidateTokenRequest,
) (*pb.ValidateTokenResponse, error) {
	if req.AccessToken == "" {
		return nil, status.Error(codes.InvalidArgument, "access token is required")
	}

	userID, err := h.service.Access.Validate(ctx, req.AccessToken)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrInvalidToken):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, jwt.ErrExpiredToken):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Error(
				codes.Internal,
				"something went wrong, try again later",
			)
		}
	}

	return &pb.ValidateTokenResponse{UserId: userID.String()}, nil
}
