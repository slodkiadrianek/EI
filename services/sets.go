package services

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/slodkiadrianek/EI/repositories"
	"github.com/slodkiadrianek/EI/utils"
)

type SetsService struct {
	SetsRepository *repositories.SetsRepository
	LoggerService  *utils.Logger
}

func NewSetsService(setsRepository *repositories.SetsRepository, loggerService *utils.Logger) *SetsService {
	return &SetsService{
		SetsRepository: setsRepository,
		LoggerService:  loggerService,
	}
}

func (s *SetsService) CreateSet(ctx context.Context, file multipart.File, name string) error {
	records, err := utils.ParseCsv(file)
	if err != nil {
		return err
	}
	fmt.Println(records)
	return nil
}
