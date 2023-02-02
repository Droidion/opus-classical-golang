package models

import (
	"github.com/rotisserie/eris"
)

type Genre struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Works []Work `json:"works"`
}

func (repo *Repo) GetGenres(composerId int) ([]*Genre, error) {
	var genres []*Genre
	sql := "SELECT genres_and_works_by_composer($1) AS json"
	genres, err := extractSql[[]*Genre](repo.Db, sql, composerId)
	if err != nil {
		return genres, eris.Wrap(err, "Could not get genres from database.")
	}
	return genres, nil
}
