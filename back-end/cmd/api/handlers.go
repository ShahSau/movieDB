package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "available",
		Message: "Welcome to the Example API",
		Version: "1.0.0",
	}
	out, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "2023-12-14")
	Family := models.Movie{
		ID:               1,
		Title:            "The Family Plan",
		Adult:            false,
		OriginalLanguage: "en",
		Description:      "Dan Morgan is many things: a devoted husband, a loving father, a celebrated car salesman. He's also a former assassin. And when his past catches up to his present, he's forced to take his unsuspecting family on a road trip unlike any other.",
		// Image            string    `json:"image"`
		ReleaseDate: rd,
		Runtime:     124,
		MediaType:   "movie",
		Popularity:  321.553,
		VoteAverage: 7.557,
		VoteCount:   183,
	}

	movies = append(movies, Family)
	rd, _ = time.Parse("2006-01-02", "2023-12-08")
	Chicken := models.Movie{
		ID:               2,
		Title:            "Chicken Run: Dawn of the Nugget",
		Adult:            false,
		OriginalLanguage: "en",
		Description:      "A band of fearless chickens flock together to save poultry-kind from an unsettling new threat: a nearby farm that's cooking up something suspicious.",
		// Image            string    `json:"image"`
		ReleaseDate: rd,
		Runtime:     134,
		MediaType:   "movie",
		Popularity:  383.553,
		VoteAverage: 7.6,
		VoteCount:   123,
	}

	movies = append(movies, Chicken)

	out, err := json.Marshal(movies)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
