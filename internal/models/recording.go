package models

import (
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rotisserie/eris"
)

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

func (r *Recording) Process() {
	r.LengthFormatted = utils.FormatWorkLength(r.Length.Int32)
	r.RecordingPeriod = utils.FormatYearsRangeString(r.YearStart.Int32, r.YearFinish.Int32)
}

func (repo *Repo) GetRecordings(workId int) ([]*Recording, error) {
	var recordings []*Recording
	sql := "SELECT recordings_by_work($1) AS json"
	recordings, err := extractSql[[]*Recording](repo.Db, sql, workId)
	if err != nil {
		return recordings, eris.Wrap(err, "Could not get recordings from database.")
	}
	return recordings, nil
}
