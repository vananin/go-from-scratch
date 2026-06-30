package db

import (
	"database/sql"
	"fmt"
	"vananin/go-from-scratch/internal/config"
)

func NewPostgres(cfg config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(cfg.MaxOpenConns)
	conn.SetMaxIdleConns(cfg.MaxIdleConns)
	conn.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	conn.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return conn, nil
}
