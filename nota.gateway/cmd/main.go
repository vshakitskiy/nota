package main

import (
	"context"
	"log"

	"nota.gateway/internal/app"
)

func main() {
	ctx := context.Background()

	app := app.NewApp()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
