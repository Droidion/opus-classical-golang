package models

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/rotisserie/eris"
)

// Genre represents musical genre, like Symphonies, or Concertante, or Choral, and all works belonging to the genre.
type Genre struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Works []Work `json:"works"`
}

// GetGenres returns all genres and music works of each genre, belonging to a given composer.
func (repo *Repo) GetGenres(composerId int) ([]*Genre, error) {
	var genres []*Genre

	sql, _, err := dialect.Select(goqu.Func("genres_and_works_by_composer", composerId).As("json")).ToSQL()
	if err != nil {
		return nil, eris.Wrapf(err, "Failed to construct SQL query with goqu for getting genres with composer ID %d", composerId)
	}

	genres, err = extractSql[[]*Genre](repo.Db, sql)
	if err != nil {
		return genres, eris.Wrapf(err, "Failed to parse JSON for composer with composer ID %d ", composerId)
	}
	return genres, nil
}
