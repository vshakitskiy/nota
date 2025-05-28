package handler

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	auth_v1 "nota.auth/pkg/pb/v1"
	"nota.shared/env"
)

func SetupHandlers(ctx context.Context, mux *runtime.ServeMux, dialOpts []grpc.DialOption) error {
	if err := auth_v1.RegisterAuthServiceHandlerFromEndpoint(
		ctx, mux, env.GetAuthHost()+":"+env.GetAuthPort(), dialOpts,
	); err != nil {
		return err
	}

	return nil
}
