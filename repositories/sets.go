package repositories

import (
	"github.com/slodkiadrianek/EI/config"
	"github.com/slodkiadrianek/EI/utils"
)

type SetsRepository struct {
	LoggerService *utils.Logger
	Db            *config.Db
}

func NewSetsRepository(loggerService *utils.Logger, db *config.Db) *SetsRepository {
	return &SetsRepository{
		LoggerService: loggerService,
		Db:            db,
	}
}
