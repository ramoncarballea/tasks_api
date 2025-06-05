package database

import (
	"database/sql"
)

type Connection interface {
	Open() (*sql.DB, error)
	ApplyMigrations() error
}
