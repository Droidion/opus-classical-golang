package models

import (
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rotisserie/eris"
	"strings"
)

type Composer struct {
	Id                int         `json:"id"`
	LastName          string      `json:"lastName"`
	FirstName         string      `json:"firstName"`
	YearBorn          pgtype.Int4 `json:"yearBorn"`
	YearDied          pgtype.Int4 `json:"yearDied"`
	YearsLived        string
	Countries         []string `json:"countries"`
	CountriesRendered string
	Slug              string      `json:"slug"`
	Enabled           pgtype.Bool `json:"enabled"`
	WikipediaLink     pgtype.Text `json:"wikipediaLink"`
	ImslpLink         pgtype.Text `json:"imslpLink"`
}

func (c *Composer) Process() {
	c.CountriesRendered = strings.Join(c.Countries, ", ")
	c.YearsLived = utils.FormatYearsRangeString(c.YearBorn.Int32, c.YearDied.Int32)
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
