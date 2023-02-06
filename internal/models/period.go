package models

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rotisserie/eris"
)

// Period represents musical period, like Baroque, or Romanticism.
type Period struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	YearStart   pgtype.Int4 `json:"yearStart"`
	YearEnd     pgtype.Int4 `json:"yearEnd"`
	YearsLasted string
	Slug        pgtype.Text `json:"slug"`
	Composers   []Composer  `json:"composers"`
}

// EnrichForTemplate adds period data needed during template render.
func (p *Period) EnrichForTemplate() {
	p.YearsLasted = fmt.Sprintf("%d–", p.YearStart.Int32)
	if p.YearEnd.Valid {
		p.YearsLasted = p.YearsLasted + fmt.Sprintf("%d", p.YearEnd.Int32)
	}
}

// GetPeriods returns periods from database.
func (repo *Repo) GetPeriods() ([]*Period, error) {
	var periods []*Period
	sql, _, err := dialect.
		From("periods_composers").
		Select("json").
		ToSQL()
	if err != nil {
		return periods, eris.Wrap(err, "Failed to construct SQL query with goqu for getting periods")
	}
	periods, err = extractSql[[]*Period](repo.Db, sql)
	if err != nil {
		return periods, eris.Wrap(err, "Failed to parse JSON for periods")
	}
	return periods, nil
}
