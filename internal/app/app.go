package app

import (
	grpcapp "fileservices/internal/app/grpc"
	"fileservices/internal/config"
	"fileservices/internal/postgres"
	"fileservices/internal/services/fileServices"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
	DB      *pgx.Conn
}

func New(log *slog.Logger, cfg *config.Config) (*App, error) {
	db, err := postgres.InitDatabase(log, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	fileService := fileServices.New(log, &postgres.Postgres{}) // Передайте экземпляр Postgres в качестве FileSaver

	grpcApp := grpcapp.New(log, fileService, cfg.GRPC.Port)

	return &App{
		GRPCSrv: grpcApp,
		DB:      db,
	}, nil
}
