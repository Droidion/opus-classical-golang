package models

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/rotisserie/eris"
)

type Genre struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Works []Work `json:"works"`
}

func (repo *Repo) GetGenres(composerId int) ([]*Genre, error) {
	var genres []*Genre

	sql, _, err := dialect.Select(goqu.Func("genres_and_works_by_composer", composerId).As("json")).ToSQL()
	if err != nil {
		return nil, eris.Wrap(err, "Could not construct SQL request to get genres from database.")
	}

	genres, err = extractSql[[]*Genre](repo.Db, sql)
	if err != nil {
		return genres, eris.Wrap(err, "Could not get genres from database.")
	}
	return genres, nil
}
