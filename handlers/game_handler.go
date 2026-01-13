package handlers

import (
	"database/sql"
	"github.com/gnujesus/arch-gaming-api/models"
)

type GameHandler struct {
	DB *sql.DB
}

func (h *GameHandler) GetAllGames() []models.Videogame {

}
