package models

type Performer struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Priority   int    `json:"priority"`
	Instrument string `json:"instrument"`
}
