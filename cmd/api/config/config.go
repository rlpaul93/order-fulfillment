
package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	APIPort     string
}

func Load() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/order_fulfillment?sslmode=disable"
	}
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{
		DatabaseURL: dbURL,
		APIPort:     port,
	}
}
