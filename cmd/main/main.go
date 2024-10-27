package main

import (
	"github.com/s4bb4t/sgRPC/internal/app"
	"github.com/s4bb4t/sgRPC/internal/config"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg)

	log.Debug("DEBUG MODE ENABLED")
	log.Info("Welcome to sgRPC")

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.GRPCSrv.MustRun()

	// TODO: run gRPC app
}

const (
	envProd  = "prod"
	envDev   = "dev"
	envLocal = "local"
)

func setupLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	}

	return log
}
