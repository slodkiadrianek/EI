package services

import (
	"context"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/repositories"
	"github.com/slodkiadrianek/EI/schema"
	"github.com/slodkiadrianek/EI/utils"
)

type CategoryService struct {
	CategoryRepository *repositories.CategoryRepository
	LoggerService      *utils.Logger
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository, loggerService *utils.Logger) *CategoryService {
	return &CategoryService{
		CategoryRepository: categoryRepository,
		LoggerService:      loggerService,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, category *schema.CreateCategory) error {
	categoryDTO := DTO.NewCategory(category.Name, category.Description)
	err := c.CategoryRepository.CreateCategory(ctx, *categoryDTO)
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryService) GetCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := c.CategoryRepository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, categoryId int) (models.Category, error){
	category , err := c.CategoryRepository.GetCategory(ctx, categoryId)
	if err != nil{
		return models.Category{}, err
	}
	return category, nil
}


func (c *CategoryService) GetCategoryWithSets(ctx context.Context, categoryId int) ([]models.CategoryWithSet, error) {
	categoriesWithSets, err := c.CategoryRepository.GetCategoryWithSets(ctx, categoryId)
	if err != nil {
		return []models.CategoryWithSet{}, err
	}
	return categoriesWithSets, nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, categoryId int) error{
	err := c.CategoryRepository.DeleteCategory(ctx, categoryId)
	if err != nil{
		return err
	}
	return nil
}