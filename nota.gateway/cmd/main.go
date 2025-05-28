package main

import (
	"context"
	"log"

	"nota.gateway/internal/app"
	"nota.shared/config"
	"nota.shared/env"
)

func main() {
	cfg, err := config.LoadGateway()
	if err != nil {
		log.Fatalf("failed to load gateway config: %v", err)
	}

	if err := env.LoadEnv(cfg.Env); err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	app := app.NewApp(cfg)

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
