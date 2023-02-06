package models

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/droidion/opus-classical-golang/internal/utils"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rotisserie/eris"
)

// Work represents music work, like Symphony No. 9 by Beethoven.
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

// EnrichForTemplate adds work data needed during template render.
func (w *Work) EnrichForTemplate() {
	w.FullName = utils.FormatWorkName(w.Title, w.No.Int32, w.Nickname.String)
	w.CatalogueNotation = utils.FormatCatalogueName(w.CatalogueName.String, w.CatalogueNumber.Int32, w.CataloguePostfix.String)
	w.ComposePeriod = utils.FormatYearsRangeString(w.YearStart.Int32, w.YearFinish.Int32)
	w.AverageLengthFormatted = utils.FormatWorkLength(w.AverageMinutes.Int32)
}

const worksTable, keysTable, cataloguesTable = "works", "keys", "catalogues"

var worksDataset = dialect.
	From(worksTable).
	LeftJoin(goqu.T(cataloguesTable), goqu.On(goqu.Ex{"works.catalogue_id": goqu.I("catalogues.id")})).
	LeftJoin(goqu.T(keysTable), goqu.On(goqu.Ex{"works.key_id": goqu.I("keys.id")})).
	Select(goqu.C("id").Table(worksTable),
		goqu.C("title").Table(worksTable),
		goqu.C("year_start").Table(worksTable),
		goqu.C("year_finish").Table(worksTable),
		goqu.C("average_minutes").Table(worksTable),
		goqu.C("name").Table(cataloguesTable).As("key"),
		goqu.C("no").Table(worksTable),
		goqu.C("nickname").Table(worksTable),
	)

// GetWork returns work by its ID.
func (repo *Repo) GetWork(id int) (*Work, error) {
	var works []*Work
	sql, _, err := worksDataset.Where(goqu.Ex{"works.id": id}).ToSQL()
	if err != nil {
		return nil, eris.Wrapf(err, "Failed to construct SQL query with goqu for getting work with work ID %d", id)
	}
	err = pgxscan.Select(context.Background(), repo.Db, &works, sql)
	if err != nil || len(works) == 0 {
		return nil, eris.Wrapf(err, "Failed to map work to Go struct with work ID %d", id)
	}
	return works[0], nil
}

// GetChildWorks returns all children works by parent work ID.
func (repo *Repo) GetChildWorks(parentWorkId int) ([]*Work, error) {
	var works []*Work

	sql, _, err := worksDataset.
		Where(goqu.Ex{"works.parent_work_id": parentWorkId}).
		Order(
			goqu.I("sort").Asc(),
			goqu.I("year_finish").Asc(),
			goqu.I("no").Asc(),
			goqu.I("catalogue_number").Asc(),
			goqu.I("catalogue_postfix").Asc(),
			goqu.I("nickname").Asc(),
		).
		ToSQL()
	if err != nil {
		return works, eris.Wrapf(err, "Failed to construct SQL query with goqu for getting children works with parent work ID %d", parentWorkId)
	}

	err = pgxscan.Select(context.Background(), repo.Db, &works, sql)
	if err != nil {
		return works, eris.Wrapf(err, "Failed to map children works to Go struct with parent work ID %d", parentWorkId)
	}
	return works, nil
}
