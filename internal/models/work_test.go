package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessWork(t *testing.T) {
	work := &Work{
		Title:            "Symphony",
		No:               pgtype.Int4{Int32: 9, Valid: true},
		Nickname:         pgtype.Text{String: "Great", Valid: true},
		CatalogueName:    pgtype.Text{String: "BWV", Valid: true},
		CatalogueNumber:  pgtype.Int4{Int32: 256, Valid: true},
		CataloguePostfix: pgtype.Text{String: "a", Valid: true},
		YearStart:        pgtype.Int4{Int32: 1900, Valid: true},
		YearFinish:       pgtype.Int4{Int32: 1970, Valid: true},
		AverageMinutes:   pgtype.Int4{Int32: 80, Valid: true},
	}
	work.EnrichForTemplate()
	assert.Equal(t, "Symphony No. 9&nbsp;<em>Great</em>", work.FullName)
	assert.Equal(t, "BWV 256a", work.CatalogueNotation)
	assert.Equal(t, "1900–70", work.ComposePeriod)
	assert.Equal(t, "1h 20m", work.AverageLengthFormatted)
}
