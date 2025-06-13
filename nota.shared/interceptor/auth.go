package interceptor

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}

		userIDValues := md.Get("x-user-id")
		if len(userIDValues) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}
		userIDStr := userIDValues[0]

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}

		ctx = context.WithValue(ctx, userIDKey, userID)

		return handler(ctx, req)
	}
}

func GetUserID(ctx context.Context) uuid.UUID {
	userID, ok := ctx.Value(userIDKey).(uuid.UUID)
	if !ok {
		panic("userID not found in context. This should not happen. Please ensure that the AuthUnaryServerInterceptor is used.")
	}

	return userID
}
