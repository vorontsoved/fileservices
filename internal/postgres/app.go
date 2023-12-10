package postgres

import (
	"context"
	"fileservices/internal/config"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type Postgres struct {
	db *pgx.Conn
}

func InitDatabase(log *slog.Logger, pgConfig config.PostgresConfig) (*pgx.Conn, error) {
	connConfig := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.Database)

	conn, err := pgx.Connect(context.Background(), connConfig)
	if err != nil {
		log.Info("Error connecting to the database: %v\n", err)
		return nil, err
	}

	log.Info("Successfully connected to the database")

	return conn, nil
}

func (p *Postgres) FileSaver(ctx context.Context, file_name string) (id int, err error) {
	// Здесь нужно добавить код сохранения файла в базу данных
	return 0, nil
}
