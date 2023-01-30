package models

import "github.com/jackc/pgx/v5/pgxpool"

type ProvidesData interface {
}

type Repo struct {
	Db *pgxpool.Pool
}

const PostgresName = "postgres"
