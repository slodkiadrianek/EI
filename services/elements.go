package services

import (
	"context"
	"mime/multipart"
	"strings"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/repositories"
	"github.com/slodkiadrianek/EI/utils"
)

type ElementService struct {
	ElementRepository *repositories.ElementReplository
	LoggerService     *utils.Logger
}

func NewElementService(elementRepository *repositories.ElementReplository, loggerService *utils.Logger) *ElementService {
	return &ElementService{
		ElementRepository: elementRepository,
		LoggerService:     loggerService,
	}
}

func (e *ElementService) CreateElements(ctx context.Context, file multipart.File, name string, setId int) error {
	records, err := utils.ParseCsv(file)
	if err != nil {
		return err
	}
	var elements []DTO.Element
	for _, record := range records {
		splittedRecord := strings.Split(strings.Join(record, ""), "/")
		if len(splittedRecord) < 4 {
			return models.NewError(400, "InvalidData", "Invalid record format: "+strings.Join(record, ","))
		}
		element := &DTO.Element{
			English:         splittedRecord[0],
			Polish:          splittedRecord[1],
			ExampleSentence: splittedRecord[2],
			Synonym:         splittedRecord[3],
			SetId:           setId,
		}
		elements = append(elements, *element)
	}
	err = e.ElementRepository.CreateNewElements(ctx, elements)
	if err != nil {
		return err
	}
	return nil
}

func (e *ElementService) GetElementsBySetId(ctx context.Context, setId int) ([]models.Element, error) {
	elements, err := e.ElementRepository.GetElementsBySetId(ctx, setId)
	if err != nil {
		return []models.Element{}, err
	}
	return elements, nil
}

