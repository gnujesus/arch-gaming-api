package services

import "github.com/gnujesus/arch-gaming-api/models"

type PetitionService interface {
	GenericService[models.Petition]
}

func NewPetitionService() *PetitionService {
	return &PetitionService{}
}
