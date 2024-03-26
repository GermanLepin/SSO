package app

import (
	grpc_app "sso/internal/app/grpc"
	"time"

	"log/slog"
)

type App struct {
	GRPCServer *grpc_app.GRPCApp
}

func NewApp(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	//storage, err := sqlite.New(storagePath)
	//

	// authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpc_app.NewGRPCApp(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
