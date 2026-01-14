package services

import (
	"database/sql"
	"github.com/gnujesus/arch-gaming-api/config"
	"github.com/gnujesus/arch-gaming-api/models"
)

type SqlPetitionService struct {
	DB     *sql.DB
	Config *config.AppConfig
}

func NewSqlPetitionService(db *sql.DB, config *config.AppConfig) *SqlPetitionService {
	return &SqlPetitionService{
		DB:     db,
		Config: config,
	}
}

func (s *SqlPetitionService) GetAll() ([]models.Petition, error) {
	return nil, nil
}

func (s *SqlPetitionService) GetById(id string) (models.Petition, error) {
	return models.Petition{}, nil
}

func (s *SqlPetitionService) Save(item models.Petition) error {
	return nil
}

func (s *SqlPetitionService) Delete(id string) error {
	return nil
}
