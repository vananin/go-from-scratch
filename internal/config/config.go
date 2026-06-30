package config

import (
	"fmt"
	"time"
)

const (
	HTTP_PORT_DEFAULT = 8080
	DB_PORT_DEFAULT   = 5432
	DB_MAX_OPEN_CONNS = 25
	DB_MAX_IDLE_CONNS = 25
)

type Config struct {
	App  AppConfig
	HTTP HTTPConfig
	DB   DBConfig
}

type AppConfig struct {
	Env string
}

type HTTPConfig struct {
	Port int
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

func Load() (Config, error) {
	httpPort, err := envInt("HTTP_PORT", HTTP_PORT_DEFAULT)
	if err != nil {
		return Config{}, fmt.Errorf("invalid HTTP_PORT: %w", err)
	}

	dbPort, err := envInt("DB_PORT", DB_PORT_DEFAULT)
	if err != nil {
		return Config{}, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	maxOpenConns, err := envInt("DB_MAX_OPEN_CONNS", DB_MAX_OPEN_CONNS)
	if err != nil {
		return Config{}, fmt.Errorf("invalid DB_MAX_OPEN_CONNS: %w", err)
	}

	maxIdleConns, err := envInt("DB_MAX_IDLE_CONNS", DB_MAX_IDLE_CONNS)
	if err != nil {
		return Config{}, fmt.Errorf("invalid DB_MAX_IDLE_CONNS: %w", err)
	}

	connMaxLifetime, err := envDuration("DB_CONN_MAX_LIFETIME", time.Hour)
	if err != nil {
		return Config{}, fmt.Errorf("invalid DB_CONN_MAX_LIFETIME: %w", err)
	}

	connMaxIdleTime, err := envDuration("DB_CONN_MAX_IDLE_TIME", 5*time.Minute)
	if err != nil {
		return Config{}, fmt.Errorf("invalid DB_CONN_MAX_IDLE_TIME: %w", err)
	}

	cfg := Config{
		App: AppConfig{
			Env: envString("APP_ENV", "development"),
		},
		HTTP: HTTPConfig{
			Port: httpPort,
		},
		DB: DBConfig{
			Host:            envString("DB_HOST", "localhost"),
			Port:            dbPort,
			User:            envString("DB_USER", "postgres"),
			Password:        envString("DB_PASSWORD", ""),
			Name:            envString("DB_NAME", "app"),
			SSLMode:         envString("DB_SSL_MODE", "enable"),
			MaxOpenConns:    maxOpenConns,
			MaxIdleConns:    maxIdleConns,
			ConnMaxLifetime: connMaxLifetime,
			ConnMaxIdleTime: connMaxIdleTime,
		},
	}

	return cfg, nil
}
