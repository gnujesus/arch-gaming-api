package services

import "github.com/gnujesus/arch-gaming-api/models"

type GameRequestService interface {
	GenericService[models.GameRequest]
}
