package handler

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	auth_v1 "nota.auth/pkg/pb/v1"
)

func SetupHandlers(ctx context.Context, mux *runtime.ServeMux, dialOpts []grpc.DialOption) error {
	if err := auth_v1.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:40401", dialOpts); err != nil {
		return err
	}

	if err := auth_v1.RegisterAccessServiceHandlerFromEndpoint(ctx, mux, "localhost:40401", dialOpts); err != nil {
		return err
	}

	return nil
}
