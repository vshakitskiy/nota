package main

import (
	"context"
	"log"

	"nota.shared/config"
	"nota.shared/database"
	"nota.shared/env"
	. "nota.snippet/internal/app"
	"nota.snippet/internal/model"
)

func main() {
	cfg, err := config.LoadSnippet()
	if err != nil {
		log.Fatalf("failed to load snippet config: %v", err)
	}

	if err := env.LoadEnv(cfg.Env); err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Running migration..")
	if err = database.Migrate(db, model.Snippet{}); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Starting app...")
	app := NewApp(db, cfg)
	ctx := context.Background()

	if err := app.Start(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
