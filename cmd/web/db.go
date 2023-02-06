package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// initDb returns new database connection.
func (app *application) initDb() *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), app.config.DatabaseConnectionString)
	if err != nil {
		app.logger.FatalWithContext("Failed to get new database pool", err, "config.DatabaseConnectionString", app.config.DatabaseConnectionString)
	}
	return pool
}
