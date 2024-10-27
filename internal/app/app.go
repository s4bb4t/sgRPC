package app

import (
	grpcapp "github.com/s4bb4t/sgRPC/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, StoragePath string, TokenTTL time.Duration) *App {
	// TODO: db

	// TODO: init grpc service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{GRPCSrv: grpcApp}
}
