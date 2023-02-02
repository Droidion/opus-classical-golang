package models

import (
	"context"
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/rotisserie/eris"
)

type Work struct {
	Id                     int    `json:"id" db:"id"`
	Title                  string `json:"title" db:"title"`
	YearStart              int    `json:"yearStart" db:"year_start"`
	YearFinish             int    `json:"yearFinish" db:"year_finish"`
	ComposePeriod          string
	AverageMinutes         int `json:"averageMinutes" db:"average_minutes"`
	AverageLengthFormatted string
	CatalogueName          string `json:"catalogueName" db:"catalogue_name"`
	CatalogueNumber        int    `json:"catalogueNumber" db:"catalogue_number"`
	CataloguePostfix       string `json:"cataloguePostfix" db:"catalogue_postfix"`
	CatalogueNotation      string
	FullName               string
	Key                    string `json:"key" db:"key"`
	No                     int    `json:"no" db:"no"`
	Nickname               string `json:"nickname" db:"nickname"`
}

func (w *Work) Process() {
	w.FullName = utils.FormatWorkName(w.Title, w.No, w.Nickname)
	w.CatalogueNotation = utils.FormatCatalogueName(w.CatalogueName, w.CatalogueNumber, w.CataloguePostfix)
	w.ComposePeriod = utils.FormatYearsRangeString(w.YearStart, w.YearFinish)
	w.AverageLengthFormatted = utils.FormatWorkLength(w.AverageMinutes)
}

func (repo *Repo) GetWork(id int) (*Work, error) {
	var works []*Work
	sql := `select w.id,
           w.title,
           COALESCE(w.year_start, 0) as year_start,
           COALESCE(w.year_finish, 0) as year_finish,
           COALESCE(w.average_minutes, 0) as average_minutes,
           COALESCE(c.name, '') as catalogue_name,
           COALESCE(w.catalogue_number, 0) as catalogue_number,
           COALESCE(w.catalogue_postfix, '') as catalogue_postfix,
           COALESCE(k.name, '') as key,
           COALESCE(w.no, 0) as no,
           COALESCE(w.nickname, '') as nickname
    from works w
             left join catalogues c on w.catalogue_id = c.id
             left join keys k on w.key_id = k.id
    where w.id = $1`
	err := pgxscan.Select(context.Background(), repo.Db, &works, sql, id)
	if err != nil || len(works) == 0 {
		return nil, eris.Wrap(err, "Could not get work from database.")
	}
	return works[0], nil
}

func (repo *Repo) GetChildWork(parentWorkId int) ([]*Work, error) {
	var works []*Work
	sql := `select w.id,
           w.title,
           COALESCE(w.year_start, 0) as year_start,
           COALESCE(w.year_finish, 0) as year_finish,
           COALESCE(w.average_minutes, 0) as average_minutes,
           COALESCE(c.name, '') as catalogue_name,
           COALESCE(w.catalogue_number, 0) as catalogue_number,
           COALESCE(w.catalogue_postfix, '') as catalogue_postfix,
           COALESCE(k.name, '') as key,
           COALESCE(w.no, 0) as no,
           COALESCE(w.nickname, '') as nickname
    from works w
             left join catalogues c on w.catalogue_id = c.id
             left join keys k on w.key_id = k.id
    where w.parent_work_id = $1
    order by sort, year_finish, no, catalogue_number, catalogue_postfix, nickname`
	err := pgxscan.Select(context.Background(), repo.Db, &works, sql, parentWorkId)
	if err != nil {
		return works, eris.Wrap(err, "Could not get works from database.")
	}
	return works, nil
}
