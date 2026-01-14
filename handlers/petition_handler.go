package handlers

import (
	"github.com/gnujesus/arch-gaming-api/interfaces/services"
	"log"
	"net/http"
)

type PetitionHandler struct {
	Service *services.PetitionService
	Logger  log.Logger
}

func NewPetitionHandler(s *services.PetitionService, l log.Logger) *PetitionHandler {
	return &PetitionHandler{
		Service: s,
		Logger:  l,
	}
}

func (h *PetitionHandler) GetAllPetitions() {

}

func (h *PetitionHandler) SubmitPetition(w http.ResponseWriter, r *http.Request) {

}
