package services

import (
	"context"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/repositories"
	"github.com/slodkiadrianek/EI/schema"
	"github.com/slodkiadrianek/EI/utils"
)

type SetsService struct {
	SetRepository *repositories.SetRepository
	ElementRepository *repositories.ElementReplository
	LoggerService *utils.Logger
}

func NewSetsService(setsRepository *repositories.SetRepository, elementRepository *repositories.ElementReplository,loggerService *utils.Logger) *SetsService {
	return &SetsService{
		SetRepository: setsRepository,
		ElementRepository: elementRepository,
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

func (s *SetsService) GetSetsWithElements(ctx context.Context) ([]models.SetWithElements, error){
	sets, err := s.SetRepository.GetSets(ctx)
	if err != nil{
		return []models.SetWithElements{} , err
	}
	elements, err := s.ElementRepository.GetElements(ctx)
	if err != nil{
		return []models.SetWithElements{}, err
	}
	setsWithElements := make([]models.SetWithElements, len(sets))
	for i := range sets {
		for _, element := range elements {
			if sets[i].Id == element.SetId {
				setsWithElements[i].CategoryId = sets[i].CategoryId
				setsWithElements[i].Description = sets[i].Description
				setsWithElements[i].Id = sets[i].Id
				setsWithElements[i].Name = sets[i].Name
				setsWithElements[i].Elements = append(setsWithElements[i].Elements, element)
			}
		}
	}
	return setsWithElements, nil
}

func (s *SetsService) DeleteSet(ctx context.Context, setId int) error{
	err := s.SetRepository.DeleteSet(ctx, setId)
	if err != nil{
		return err
	}
	return nil
}