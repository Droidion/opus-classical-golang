package models

type Composer struct {
	Id            int      `json:"id"`
	LastName      string   `json:"lastName"`
	FirstName     string   `json:"firstName"`
	YearBorn      int      `json:"yearBorn"`
	YearDied      int      `json:"yearDied"`
	Countries     []string `json:"countries"`
	Slug          string   `json:"slug"`
	Enabled       bool     `json:"enabled"`
	WikipediaLink string   `json:"wikipediaLink"`
	ImslpLink     string   `json:"imslpLink"`
}
