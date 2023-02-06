package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessComposer(t *testing.T) {
	composer := &Composer{
		Countries: []string{"Germany", "France"},
		YearBorn:  pgtype.Int4{Int32: 1900, Valid: true},
		YearDied:  pgtype.Int4{Int32: 1970, Valid: true},
	}
	composer.EnrichForTemplate()
	assert.Equal(t, "Germany, France", composer.CountriesRendered)
	assert.Equal(t, "1900–70", composer.YearsLived)
}
