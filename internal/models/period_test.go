package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessPeriodFullYear(t *testing.T) {
	period := &Period{
		YearStart: pgtype.Int4{Int32: 1900, Valid: true},
		YearEnd:   pgtype.Int4{Int32: 1970, Valid: true},
	}
	period.EnrichForTemplate()
	assert.Equal(t, "1900–1970", period.YearsLasted)
}

func TestProcessPeriodOnlyStartYear(t *testing.T) {
	period := &Period{
		YearStart: pgtype.Int4{Int32: 1900, Valid: true},
		YearEnd:   pgtype.Int4{Int32: 0, Valid: false},
	}
	period.EnrichForTemplate()
	assert.Equal(t, "1900–", period.YearsLasted)
}
