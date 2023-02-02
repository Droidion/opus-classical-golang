package models

import (
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/rotisserie/eris"
	"strings"
)

type Composer struct {
	Id                int    `json:"id"`
	LastName          string `json:"lastName"`
	FirstName         string `json:"firstName"`
	YearBorn          int    `json:"yearBorn"`
	YearDied          int    `json:"yearDied"`
	YearsLived        string
	Countries         []string `json:"countries"`
	CountriesRendered string
	Slug              string `json:"slug"`
	Enabled           bool   `json:"enabled"`
	WikipediaLink     string `json:"wikipediaLink"`
	ImslpLink         string `json:"imslpLink"`
}

func (c *Composer) Process() {
	c.CountriesRendered = strings.Join(c.Countries, ", ")
	c.YearsLived = utils.FormatYearsRangeString(c.YearBorn, c.YearDied)
}

func (repo *Repo) GetComposer(slug string) (*Composer, error) {
	var composer *Composer
	sql := "SELECT composer_by_slug($1) AS json"
	composer, err := extractSql[*Composer](repo.Db, sql, slug)
	if err != nil {
		return nil, eris.Wrap(err, "Could not get composer from database.")
	}
	return composer, nil
}
