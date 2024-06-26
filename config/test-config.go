package config

import (
	"context"
	"database/sql"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ctx context.Context
	connString string
)


func setupContainer() {
	ctx = context.Background()
	c, err := postgres.RunContainer(
		ctx,
		testcontainers.WithImage("postgres:14-alpine"),
		postgres.WithDatabase("test"),
		postgres.WithUsername("postgre"),
		postgres.WithPassword("postgre"),
		postgres.WithInitScripts("../../../config/init_db.sql"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),	
		),
	)

	if err != nil {
		panic(err)
	}
	
	connString, err = c.ConnectionString(ctx)

	connString = connString + "sslmode=disable"

	if err != nil {
		panic(err)
	}
}

func SetupDbConnection() *sql.DB {
	setupContainer()
	return InitConfig(connString)
}