package models

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/rotisserie/eris"
)

type SearchResult struct {
	Id            int     `json:"id" db:"id"`
	FirstName     string  `json:"firstName" db:"first_name"`
	LastName      string  `json:"lastName" db:"last_name"`
	Slug          string  `json:"slug" db:"slug"`
	LastNameScore float64 `json:"lastNameScore" db:"last_name_score"`
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
