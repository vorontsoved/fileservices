package app

import (
	grpcapp "fileservices/internal/app/grpc"
	"fileservices/internal/config"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, cfg *config.Config) (*App, error) {
	//TODO: инициализировать хранилище (storage)

	db, err := initDatabase(log, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	grpcApp := grpcapp.New(log, cfg.GRPC.Port)

	return &App{
		GRPCSrv: grpcApp,
	}, nil
}
