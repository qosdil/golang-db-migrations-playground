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
	"github.com/spf13/viper"
)

const version string = "1.3.0"

// Movie ...
type Movie struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Stars       string `json:"stars,omitempty"`
	ReleaseYear string `json:"release_year,omitempty"`
	Language    string `json:"language,omitempty"`
}

// GetMovies ...
func GetMovies(w http.ResponseWriter, r *http.Request) {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	dbUsername := viper.GetString("database.username")
	dbPassword := viper.GetString("database.password")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := viper.GetString("database.name")
	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		fmt.Fprintf(w, "\n\nAn error occured during your MySQL command: %s", err)
		panic(err)
	}

	rows, err := db.Query("SELECT id, title, stars, release_year, language FROM movies ORDER BY id")
	if err != nil {
		http.Redirect(w, r, "/error", 302)
		return
	}

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var movies []Movie
	stringLanguage := ""
	stringStars := ""
	for rows.Next() {
		var (
			id           string
			title        string
			stars        sql.NullString
			release_year string
			language     sql.NullString
		)
		if err := rows.Scan(&id, &title, &stars, &release_year, &language); err != nil {
			panic(err)
		}

		stringLanguage = ""
		if language.String != "" {
			stringLanguage = language.String
		}

		stringStars = ""
		if stars.String != "" {
			stringStars = stars.String
		}
		movies = append(movies, Movie{ID: id, Title: title, Stars: stringStars, ReleaseYear: release_year, Language: stringLanguage})
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	w.Header().Set("Version", version)
	json.NewEncoder(w).Encode(movies)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", GetMovies).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
