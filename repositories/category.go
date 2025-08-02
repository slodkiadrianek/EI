package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/utils"
)

type CategoryRepository struct {
	LoggerService *utils.Logger
	Db            *sql.DB
}

func NewCategoryRepository(loggerService *utils.Logger, db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		LoggerService: loggerService,
		Db:            db,
	}
}

func (c *CategoryRepository) CreateCategory(ctx context.Context, category DTO.Category) error {
	fmt.Print("CategoryRepository: CreateCategory called with category: ", category)
	query := "INSERT INTO categories(name,description) VALUES($1,$2)"
	stmt, err := c.Db.PrepareContext(ctx,query)
	if err != nil {
		c.LoggerService.Error("Failed to prepare query for execution")
		return models.NewError(500, "Database", "Failed to insert data to a database")
	}
	defer stmt.Close()
	_,err = stmt.ExecContext(ctx, category.Name, category.Description)
	if err != nil {
		c.LoggerService.Error("Failed to execute query")
		return models.NewError(500, "Database", "Failed to insert data to a database")
	}
	return nil
}

func (c *CategoryRepository) GetCategories(ctx context.Context) ([]models.Category, error) {
	query := "SELECT * FROM categories"
	stmt, err := c.Db.PrepareContext(ctx, query)
	if err != nil {
		c.LoggerService.Error("Failed to prepare query for execution")
		return []models.Category{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var categories []models.Category
	sql, err := stmt.QueryContext(ctx)
	if err != nil {
		c.LoggerService.Error("Failed to execute query")
		return []models.Category{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	for sql.Next() {
		var category models.Category
		err = sql.Scan(&category.Id, &category.Name, &category.Description)
		if err != nil {
			c.LoggerService.Error("Failed to scan row")
			return []models.Category{}, models.NewError(500, "Database", "Failed to get data from a database")
		}
		categories = append(categories, category)
	}
	if err = sql.Err(); err != nil {
		c.LoggerService.Error("Failed to iterate over rows")
		return []models.Category{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return categories, nil
}

func (c *CategoryRepository) GetCategory(ctx context.Context, categoryId int) (models.Category, error) {
	query := "SELECT * FROM categories WHERE id = $1"
	stmt, err := c.Db.PrepareContext(ctx, query)
	if err != nil {
		c.LoggerService.Error("Failed to prepare query for execution")
		return models.Category{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	defer stmt.Close()
	var data models.Category
	err = stmt.QueryRowContext(ctx, categoryId).Scan(&data.Id, &data.Name, &data.Description)
	if err == sql.ErrNoRows {
		c.LoggerService.Error("No category found with the given ID")
		return models.Category{}, models.NewError(404, "NotFound", "Category not found")
	}
	if err != nil {
		c.LoggerService.Error("Failed to execute query")
		return models.Category{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return data, nil
}
