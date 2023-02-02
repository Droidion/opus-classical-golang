package models

import (
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/rotisserie/eris"
)

type Recording struct {
	Id              int    `json:"int"`
	Label           string `json:"label"`
	Length          int    `json:"length"`
	LengthFormatted string
	CoverName       string      `json:"coverName"`
	Streamers       []Streamer  `json:"streamers"`
	YearStart       int         `json:"yearStart"`
	Performers      []Performer `json:"performers"`
	YearFinish      int         `json:"yearFinish"`
	RecordingPeriod string
}

func (r *Recording) Process() {
	r.LengthFormatted = utils.FormatWorkLength(r.Length)
	r.RecordingPeriod = utils.FormatYearsRangeString(r.YearStart, r.YearFinish)
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
