package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitConfig(connString string) (*sql.DB) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	return db
}