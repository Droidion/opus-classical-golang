package models

import (
	"github.com/doug-martin/goqu/v9"
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

	sql, _, err := dialect.Select(goqu.Func("recordings_by_work", workId).As("json")).ToSQL()
	if err != nil {
		return nil, eris.Wrap(err, "Could not construct SQL request to get recordings from database.")
	}

	recordings, err = extractSql[[]*Recording](repo.Db, sql)
	if err != nil {
		return recordings, eris.Wrap(err, "Could not get recordings from database.")
	}
	return recordings, nil
}
