package config

import (
	"os"
)

// Config holds the application configuration.
type Config struct {
	DBDriver string
	DSN      string
	Port     string
}

// Load loads the configuration from environment variables or uses default values.
func Load() *Config {
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "sqlite"
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "tasks.db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DBDriver: dbDriver,
		DSN:      dsn,
		Port:     port,
	}
}
