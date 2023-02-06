package models

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rotisserie/eris"
)

// Recording represents musical recording.
type Recording struct {
	Id              int         `json:"int"`
	Label           string      `json:"label"`
	Length          pgtype.Int4 `json:"length"`
	LengthFormatted string
	CoverName       string      `json:"coverName"`
	Streamers       []Streamer  `json:"streamers"`
	YearStart       pgtype.Int4 `json:"yearStart"`
	Performers      []Performer `json:"performers"`
	YearFinish      pgtype.Int4 `json:"yearFinish"`
	RecordingPeriod string
}

// EnrichForTemplate adds Recording data needed during template render.
func (r *Recording) EnrichForTemplate() {
	r.LengthFormatted = utils.FormatWorkLength(r.Length.Int32)
	r.RecordingPeriod = utils.FormatYearsRangeString(r.YearStart.Int32, r.YearFinish.Int32)
}

// GetRecordings returns recordings from database for a given musical work.
func (repo *Repo) GetRecordings(workId int) ([]*Recording, error) {
	var recordings []*Recording

	sql, _, err := dialect.Select(goqu.Func("recordings_by_work", workId).As("json")).ToSQL()
	if err != nil {
		return nil, eris.Wrapf(err, "Failed to construct SQL query with goqu for getting recordings with work ID %d", workId)
	}

	recordings, err = extractSql[[]*Recording](repo.Db, sql)
	if err != nil {
		return recordings, eris.Wrapf(err, "Failed to parse JSON for recordings with work ID %d ", workId)
	}
	return recordings, nil
}
