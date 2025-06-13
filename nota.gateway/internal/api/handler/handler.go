package handler

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	authpb "nota.auth/pkg/pb/v1"
	"nota.shared/env"
	snippetpb "nota.snippet/pkg/pb/v1"
)

func SetupHandlers(ctx context.Context, mux *runtime.ServeMux, dialOpts []grpc.DialOption) error {
	if err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		ctx, mux, env.GetAuthHost()+":"+env.GetAuthPort(), dialOpts,
	); err != nil {
		return err
	}

	if err := snippetpb.RegisterSnippetServiceHandlerFromEndpoint(
		ctx, mux, env.GetSnippetHost()+":"+env.GetSnippetPort(), dialOpts,
	); err != nil {
		return err
	}

	return nil
}
