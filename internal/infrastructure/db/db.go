package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// NewConnection opens a new PostgreSQL database connection.
func NewConnection(databaseURL string) (*sql.DB, error) {
	return sql.Open("postgres", databaseURL)
}
