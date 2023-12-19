package models

import (
	"time"
)

type Movie struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Adult            bool      `json:"adult"`
	OriginalLanguage string    `json:"original_language"`
	Description      string    `json:"description"`
	Image            string    `json:"image"`
	ReleaseDate      time.Time `json:"release_date"`
	Runtime          int       `json:"runtime"`
	MediaType        string    `json:"media_type"`
	Popularity       float64   `json:"popularity"`
	VoteAverage      float64   `json:"vote_average"`
	VoteCount        int       `json:"vote_count"`
	CreatedAt        time.Time `json:"-"` // ignore this field when sending data to client
	UpdatedAt        time.Time `json:"-"` // ignore this field when sending data to client
}
