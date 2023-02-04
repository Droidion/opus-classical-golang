package models

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rotisserie/eris"
)

type SearchResult struct {
	Id            int           `json:"id" db:"id"`
	FirstName     pgtype.Text   `json:"firstName" db:"first_name"`
	LastName      pgtype.Text   `json:"lastName" db:"last_name"`
	Slug          pgtype.Text   `json:"slug" db:"slug"`
	LastNameScore pgtype.Float8 `json:"lastNameScore" db:"last_name_score"`
}

func (repo *Repo) SearchComposers(query string, limit int) ([]*SearchResult, error) {
	var results []*SearchResult
	sql := `SELECT id, first_name, last_name, slug, last_name_score FROM search_composers_by_last_name($1, $2)`
	err := pgxscan.Select(context.Background(), repo.Db, &results, sql, query, limit)
	if err != nil {
		return results, eris.Wrap(err, "Could not get composers search results from database.")
	}
	return results, nil
}
