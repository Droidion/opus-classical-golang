package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessRecording(t *testing.T) {
	recording := &Recording{
		YearStart:  pgtype.Int4{Int32: 1900, Valid: true},
		YearFinish: pgtype.Int4{Int32: 1970, Valid: true},
		Length:     pgtype.Int4{Int32: 80, Valid: true},
	}
	recording.Process()
	assert.Equal(t, "1h 20m", recording.LengthFormatted)
	assert.Equal(t, "1900–70", recording.RecordingPeriod)
}
