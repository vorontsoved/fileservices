package postgres

import (
	"context"
	"fileservices/internal/config"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type Postgres struct {
	log *slog.Logger
	db  *pgx.Conn
}

func InitDatabase(log *slog.Logger, pgConfig config.PostgresConfig) (*Postgres, error) {
	connConfig := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.Database)

	conn, err := pgx.Connect(context.Background(), connConfig)
	if err != nil {
		log.Error("Error connecting to the database: %v", err)
		return nil, err
	}

	log.Info("Successfully connected to the database")

	return &Postgres{log: log, db: conn}, nil
}

func (p *Postgres) FileSaver(ctx context.Context, fileName string) (id int, err error) {
	query := "INSERT INTO files (filename, created_at, modified_at) VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id"

	var insertedID int
	err = p.db.QueryRow(ctx, query, fileName).Scan(&insertedID)
	if err != nil {
		p.log.Error("Error executing INSERT query: %v", err)
		return 0, err
	}

	p.log.Info("File saved successfully. Inserted ID: %d", insertedID)
	return insertedID, nil
}

func (p *Postgres) Close(ctx context.Context) error {
	p.log.Info("Stopping postgres")
	err := p.db.Close(ctx)
	if err != nil {
		p.log.Error("Error closing the database connection: %v", err)
	}
	return err
}
