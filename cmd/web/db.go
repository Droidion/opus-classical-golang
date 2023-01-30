package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// initDb returns new database connection.
func (app *application) initDb() *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), app.config.DatabaseConnectionString)
	if err != nil {
		app.logger.Fatal("Could not establish DB connection.", err)
	}
	return pool
}
