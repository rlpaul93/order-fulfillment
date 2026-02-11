package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	APIPort     string
	StorageMode string // "memory" (default) or "postgres"
	SwaggerHost string // Host for Swagger UI (default: remote production URL)
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
	storageMode := os.Getenv("STORAGE_MODE")
	if storageMode == "" {
		storageMode = "memory"
	}
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "order-fulfillment.onrender.com"
	}
	return &Config{
		DatabaseURL: dbURL,
		APIPort:     port,
		StorageMode: storageMode,
		SwaggerHost: swaggerHost,
	}
}
