package db

import (
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewClient(cfg Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.Database)

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
