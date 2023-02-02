package models

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/rotisserie/eris"
)

// Period represents musical period, like Baroque or Romanticism.
type Period struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	YearStart   int    `json:"yearStart"`
	YearEnd     int    `json:"yearEnd"`
	YearsLasted string
	Slug        string     `json:"slug"`
	Composers   []Composer `json:"composers"`
}

func (p *Period) Process() {
	p.YearsLasted = fmt.Sprintf("%d–", p.YearStart)
	if p.YearEnd > 0 {
		p.YearsLasted = p.YearsLasted + fmt.Sprintf("%d", p.YearEnd)
	}
}

// GetPeriods returns musical periods from database.
func (repo *Repo) GetPeriods() ([]*Period, error) {
	var periods []*Period
	sql, _, err := goqu.
		Dialect(PostgresName).
		From("periods_composers").
		Select("json").
		ToSQL()
	if err != nil {
		return periods, eris.Wrap(err, "Could not construct SQL request to get periods from database.")
	}
	periods, err = extractSql[[]*Period](repo.Db, sql)
	if err != nil {
		return periods, eris.Wrap(err, "Could not get periods from database.")
	}
	return periods, nil
}