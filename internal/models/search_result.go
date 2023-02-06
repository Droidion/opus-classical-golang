package models

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rotisserie/eris"
)

// SearchResult represents compose search results.
type SearchResult struct {
	Id            int           `json:"id" db:"id"`
	FirstName     pgtype.Text   `json:"firstName" db:"first_name"`
	LastName      pgtype.Text   `json:"lastName" db:"last_name"`
	Slug          pgtype.Text   `json:"slug" db:"slug"`
	LastNameScore pgtype.Float8 `json:"lastNameScore" db:"last_name_score"` // Postgres trigram score, from 0 to 1.
}

// SearchComposers returns compose search results.
func (repo *Repo) SearchComposers(query string, limit int) ([]*SearchResult, error) {
	var results []*SearchResult

	sql, _, err := dialect.
		From(goqu.Func("search_composers_by_last_name", query, limit)).
		Select("id", "first_name", "last_name", "slug", "last_name_score").
		ToSQL()
	if err != nil {
		return nil, eris.Wrapf(err, "Failed to construct SQL query with goqu for getting search results with query %s and limit %d", query, limit)
	}

	err = pgxscan.Select(context.Background(), repo.Db, &results, sql)
	if err != nil {
		return results, eris.Wrapf(err, "Failed to map search results to Go struct with query %s and limit %d", query, limit)
	}
	return results, nil
}
