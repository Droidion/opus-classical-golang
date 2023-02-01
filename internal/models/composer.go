package models

import (
	"github.com/droidion/opus-classical-golang/internal/utils"
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
