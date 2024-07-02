package config

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

func InitConfig(connString string) (*sql.DB) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		slog.Error("database configuration", "method", "InitConfig", "error", err)
		panic(err)
	}

	return db
}