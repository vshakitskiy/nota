package api

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nota.auth/internal/interceptor"
	"nota.auth/internal/service"
	"nota.auth/pkg/jwt"
	pb "nota.auth/pkg/pb/v1"
	"nota.shared/telemetry"
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
	ctx, span := telemetry.StartSpan(ctx, "AccessHandler.ValidateToken")
	defer span.End()

	accessToken := interceptor.GetAccessTokenStr(ctx)

	userID, err := h.service.Access.Validate(ctx, accessToken)
	if err != nil {
		telemetry.RecordError(span, err)

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
