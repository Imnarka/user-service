package main

import (
	"context"
	"github.com/Imnarka/user-service/internal/config"
	"github.com/Imnarka/user-service/internal/di"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	components, err := di.InitializeGRPCServer(cfg)
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}

	if err := components.App.Start(context.Background()); err != nil {
		components.Logger.WithError(err).Error("Application failed")
		log.Fatalf("application failed: %v", err)
	}
}
