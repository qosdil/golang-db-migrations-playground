package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Movie ...
type Movie struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	ReleaseYear string `json:"release_year,omitempty"`
}

// GetMovies ...
func GetMovies(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@/pressly_goose_test")
	if err != nil {
		fmt.Fprintf(w, "\n\nAn error occured during your MySQL command: %s", err)
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM movies ORDER BY id")
	if err != nil {
		http.Redirect(w, r, "/error", 302)
		return
	}

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var (
			id           string
			title        string
			release_year string
		)
		if err := rows.Scan(&id, &title, &release_year); err != nil {
			panic(err)
		}
		movies = append(movies, Movie{ID: id, Title: title, ReleaseYear: release_year})
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(movies)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", GetMovies).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
