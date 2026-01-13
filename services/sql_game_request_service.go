package services

import (
	"database/sql"
	"github.com/gnujesus/arch-gaming-api/config"
	"github.com/gnujesus/arch-gaming-api/models"
)

type SqlGameRequestService struct {
	DB     *sql.DB
	Config *config.AppConfig
}

func (s *SqlGameRequestService) GetAll() ([]models.GameRequest, error) {
	return nil, nil
}

func (s *SqlGameRequestService) GetById(id string) (models.GameRequest, error) {
	return models.GameRequest{}, nil
}

func (s *SqlGameRequestService) Save(item models.GameRequest) error {
	return nil
}

func (s *SqlGameRequestService) Delete(id string) error {
	return nil
}
