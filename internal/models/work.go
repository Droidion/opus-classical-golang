package models

import (
	"context"
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rotisserie/eris"
)

type Work struct {
	Id                     int         `json:"id" db:"id"`
	Title                  string      `json:"title" db:"title"`
	YearStart              pgtype.Int4 `json:"yearStart" db:"year_start"`
	YearFinish             pgtype.Int4 `json:"yearFinish" db:"year_finish"`
	ComposePeriod          string
	AverageMinutes         pgtype.Int4 `json:"averageMinutes" db:"average_minutes"`
	AverageLengthFormatted string
	CatalogueName          pgtype.Text `json:"catalogueName" db:"catalogue_name"`
	CatalogueNumber        pgtype.Int4 `json:"catalogueNumber" db:"catalogue_number"`
	CataloguePostfix       pgtype.Text `json:"cataloguePostfix" db:"catalogue_postfix"`
	CatalogueNotation      string
	FullName               string
	Key                    pgtype.Text `json:"key" db:"key"`
	No                     pgtype.Int4 `json:"no" db:"no"`
	Nickname               pgtype.Text `json:"nickname" db:"nickname"`
}

func (w *Work) Process() {
	w.FullName = utils.FormatWorkName(w.Title, w.No.Int32, w.Nickname.String)
	w.CatalogueNotation = utils.FormatCatalogueName(w.CatalogueName.String, w.CatalogueNumber.Int32, w.CataloguePostfix.String)
	w.ComposePeriod = utils.FormatYearsRangeString(w.YearStart.Int32, w.YearFinish.Int32)
	w.AverageLengthFormatted = utils.FormatWorkLength(w.AverageMinutes.Int32)
}

func (repo *Repo) GetWork(id int) (*Work, error) {
	var works []*Work
	sql := `select w.id,
           w.title,
           w.year_start,
           w.year_finish,
           w.average_minutes,
           c.name as catalogue_name,
           w.catalogue_number,
           w.catalogue_postfix,
           k.name as key,
           w.no,
           w.nickname
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
           w.year_start,
           w.year_finish,
           w.average_minutes,
           c.name as catalogue_name,
           w.catalogue_number,
           w.catalogue_postfix,
           k.name as key,
           w.no,
           w.nickname
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
