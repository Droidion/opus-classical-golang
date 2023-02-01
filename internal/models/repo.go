package models

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rotisserie/eris"
)

type ProvidesData interface {
	GetPeriods() ([]*Period, error)
}

type Repo struct {
	Db *pgxpool.Pool
}

const PostgresName = "postgres"

func extractSql[T any](db *pgxpool.Pool, sql string) (T, error) {
	var result T
	var rawJson string
	err := db.QueryRow(context.Background(), sql).Scan(&rawJson)
	if err != nil {
		return result, eris.Wrap(err, "Could not get JSON data from database.")
	}
	err = json.Unmarshal([]byte(rawJson), &result)
	if err != nil {
		return result, eris.Wrap(err, "Could not parse JSON from database.")
	}
	return result, nil
}
