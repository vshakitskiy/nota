package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"nota.auth/internal/api"
	"nota.auth/internal/repository"
	"nota.auth/internal/service"
	pb "nota.auth/pkg/pb/v1"
	"nota.shared/config"
	"nota.shared/telemetry"
)

type App struct {
	db     *gorm.DB
	port   string
	server *grpc.Server
}

func NewApp(db *gorm.DB) *App {
	return &App{
		db:   db,
		port: config.GetenvDefault("GRPC_PORT", "40401"),
	}
}

func (a *App) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	shutdownTracer, err := telemetry.NewTracerProvider(ctx, "nota.auth", "localhost:4317")
	if err != nil {
		log.Fatalf("failed to create tracer provider: %v", err)
	}

	defer func() {
		if err := shutdownTracer(ctx); err != nil {
			log.Printf("failed to shutdown tracer provider: %v", err)
		}
	}()

	a.server = grpc.NewServer(telemetry.NewGRPCServerHandlers()...)

	repo := repository.NewRepository(a.db)
	service := service.NewService(repo)

	authHandler := api.NewAuthServiceHandler(service)
	pb.RegisterAuthServiceServer(a.server, authHandler)

	accessHandler := api.NewAccessServiceHandler(service)
	pb.RegisterAccessServiceServer(a.server, accessHandler)

	reflection.Register(a.server)

	go func() {
		lis, err := net.Listen("tcp", ":"+a.port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Printf("Starting gRPC server on port %s", a.port)
		if err := a.server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return a.GracefulShutdown(ctx)
}

func (a *App) GracefulShutdown(ctx context.Context) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("Shutdown requested via context")
	case <-quit:
		log.Println("Shutdown requested via signal")
	}

	log.Println("Shutting down gRPC server...")
	a.server.GracefulStop()
	log.Println("Server gracefully stopped")

	return nil
}
