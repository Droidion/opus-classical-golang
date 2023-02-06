package models

// Performer represents musical artist, like Herbert von Karajan.
type Performer struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Priority   int    `json:"priority"`
	Instrument string `json:"instrument"`
}
