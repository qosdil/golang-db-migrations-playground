package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Movie ...
type Movie struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

// GetMovies ...
func GetMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

var movies []Movie

func main() {
	movies = append(movies, Movie{ID: "1", Title: "Rambo IV"})

	router := mux.NewRouter()
	router.HandleFunc("/movies", GetMovies).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
