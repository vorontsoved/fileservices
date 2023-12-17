package app

import (
	grpcapp "fileservices/internal/app/grpc"
	"fileservices/internal/config"
	"fileservices/internal/postgres"
	"fileservices/internal/services/fileServices"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
	DB      *postgres.Postgres
}

func New(log *slog.Logger, cfg *config.Config) (*App, error) {
	db, err := postgres.InitDatabase(log, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	fileService := fileServices.New(log, db)
	grpcApp := grpcapp.New(log, fileService, cfg.GRPC.Port)

	return &App{
		GRPCSrv: grpcApp,
		DB:      db,
	}, nil
}
