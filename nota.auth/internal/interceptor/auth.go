package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const accessTokenKey contextKey = "accessToken"

func AuthUnaryServerInterceptor(
	protectedSuffixes []string,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		isProtected := false
		for _, suffix := range protectedSuffixes {
			if strings.HasSuffix(info.FullMethod, suffix) {
				isProtected = true
				break
			}
		}

		if !isProtected {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}

		accessTokenValues := md.Get("x-access-token")
		if len(accessTokenValues) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}
		accessTokenStr := accessTokenValues[0]

		ctx = context.WithValue(ctx, accessTokenKey, accessTokenStr)

		return handler(ctx, req)
	}
}

func GetAccessTokenStr(ctx context.Context) string {
	accessToken := ctx.Value(accessTokenKey).(string)
	return accessToken
}
