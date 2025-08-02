package services

import (
	"context"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/repositories"
	"github.com/slodkiadrianek/EI/schema"
	"github.com/slodkiadrianek/EI/utils"
)

type SetsService struct {
	SetRepository *repositories.SetRepository
	LoggerService *utils.Logger
}

func NewSetsService(setsRepository *repositories.SetRepository, loggerService *utils.Logger) *SetsService {
	return &SetsService{
		SetRepository: setsRepository,
		LoggerService: loggerService,
	}
}

func (s *SetsService) CreateSet(ctx context.Context,categoryId int, set *schema.CreateSet ) (int, error) {
	setDTO := DTO.NewSet( set.Name, set.Description, categoryId)
	setId, err := s.SetRepository.CreateNewSet(ctx, setDTO)
	if err != nil {
		return 0,err
	}
	s.LoggerService.Info("Set created successfully")
	return setId,nil
}
