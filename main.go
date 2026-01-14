package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gnujesus/arch-gaming-api/config"
	"github.com/gnujesus/arch-gaming-api/handlers"
	"github.com/gnujesus/arch-gaming-api/services"
)

func setupDatabase(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	logger := log.Logger()

	config := config.Load()

	db, err := setupDatabase(config.DatabaseURL)

	petitionService := services.SqlPetitionService()

	petitionHandler := handlers.NewPetitionHandler()

	// --- Routes

	mux := http.NewServeMux()

	mux.HandleFunc("GET /games/petitions", requestHandler.SubmitGamePetition)
	mux.HandleFunc("POST /games/petitions", requestHandler.SubmitGamePetition)

	// --- Serve

	http.ListenAndServe(":3030", mux)
}
