package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

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
	logger := log.New(os.Stdout, "Log: ", log.Flags())

	config := config.Load()

	db, err := setupDatabase(config.DatabaseURL)

	if err != nil {
		log.Fatal("Database setup error")
	}

	sqlPetitionService := services.NewSqlPetitionService(db, &config)

	petitionHandler := handlers.NewPetitionHandler(sqlPetitionService, logger)

	// --- Routes

	mux := http.NewServeMux()

	mux.HandleFunc("GET /games/petitions", petitionHandler.GetAllPetitions)
	mux.HandleFunc("POST /games/petitions", petitionHandler.SubmitPetition)

	// --- Serve

	http.ListenAndServe(":3030", mux)
}
