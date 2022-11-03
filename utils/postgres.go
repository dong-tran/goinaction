package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection() (*sqlx.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	if len(host) < 1 {
		host = "localhost"
	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&application_name=%s",
		"postgres",
		"postgres",
		host,
		"25432",
		"boleto",
		"disable",
		"dummy")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db := sqlx.NewDb(conn, "postgres")
	return db, nil
}
