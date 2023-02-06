package models

import (
	"context"
	"encoding/json"
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rotisserie/eris"
)

type ProvidesData interface {
	GetPeriods() ([]*Period, error)
	GetComposer(slug string) (*Composer, error)
	GetGenres(composerId int) ([]*Genre, error)
	GetRecordings(workId int) ([]*Recording, error)
	SearchComposers(query string, limit int) ([]*SearchResult, error)
	GetWork(id int) (*Work, error)
	GetChildWork(parentWorkId int) ([]*Work, error)
}

type Repo struct {
	Db *pgxpool.Pool
}

var dialect = goqu.Dialect("postgres")

func extractSql[T any](db *pgxpool.Pool, sql string, params ...any) (T, error) {
	var result T
	var rawJson string
	var err error
	if len(params) > 0 {
		err = db.QueryRow(context.Background(), sql, params...).Scan(&rawJson)
	} else {
		err = db.QueryRow(context.Background(), sql).Scan(&rawJson)
	}
	if err != nil {
		return result, eris.Wrap(err, "Could not get JSON data from database.")
	}
	err = json.Unmarshal([]byte(rawJson), &result)
	if err != nil {
		return result, eris.Wrap(err, "Could not parse JSON from database.")
	}
	return result, nil
}
