package main

import (
	"context"
	"fileservices/internal/app"
	"fileservices/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("cfg", cfg))

	application, err := app.New(log, cfg)
	if err != nil {
		log.Info("failed init app: %v", err)
	}

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop //стоим тут пока не выполнитя сигнал

	log.Info("stopping gRPC applicstion", slog.String("sigmal", sign.String()))

	application.GRPCSrv.Stop()

	log.Info("stopping postgres", slog.String("sigmal", sign.String()))
	application.DB.Close(context.Background())

	log.Info("gRPC application stop")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
