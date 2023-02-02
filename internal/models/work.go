package models

import (
	"github.com/droidion/opus-classical-golang/internal/utils"
)

type Work struct {
	Id                     int    `json:"id"`
	Title                  string `json:"title"`
	YearStart              int    `json:"yearStart"`
	YearFinish             int    `json:"yearFinish"`
	ComposePeriod          string
	AverageMinutes         int `json:"averageMinutes"`
	AverageLengthFormatted string
	CatalogueName          string `json:"catalogueName"`
	CatalogueNumber        int    `json:"catalogueNumber"`
	CataloguePostfix       string `json:"cataloguePostfix"`
	CatalogueNotation      string
	FullName               string
	Key                    string `json:"key"`
	No                     int    `json:"no"`
	Nickname               string `json:"nickname"`
}

func (w *Work) Process() {
	w.FullName = utils.FormatWorkName(w.Title, w.No, w.Nickname)
	w.CatalogueNotation = utils.FormatCatalogueName(w.CatalogueName, w.CatalogueNumber, w.CataloguePostfix)
	w.ComposePeriod = utils.FormatYearsRangeString(w.YearStart, w.YearFinish)
	w.AverageLengthFormatted = utils.FormatWorkLength(w.AverageMinutes)
}
