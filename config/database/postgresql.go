package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"tasks.com/config/environment"
)

type postgresqlConnection struct {
	Connection
	config *environment.DataBaseConfig
}

func newPostgresqlDatabaseConnection(config *environment.DataBaseConfig) Connection {
	return &postgresqlConnection{
		config: config,
	}
}

func (con *postgresqlConnection) Open() (*sql.DB, error) {
	db, err := sql.Open("postgres", con.config.ConnectionString)
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		return nil, err
	}

	return db, nil
}

func (con *postgresqlConnection) ApplyMigrations() error {
	db, err := con.Open()
	if err != nil {
		return err
	}

	if err := goose.Up(db, "./config/database/migrations"); err != nil {
		return err
	}

	return nil
}
