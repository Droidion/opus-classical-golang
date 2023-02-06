package models

// Streamer represents a streaming service link, like for Spotify or Qobuz.
type Streamer struct {
	Icon   string `json:"icon"`
	Link   string `json:"link"`
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}
