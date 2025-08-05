package repositories

import (
	"context"
	"database/sql"
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

func (c *CategoryRepository) GetCategoryWithSets(ctx context.Context, categoryId int) ([]models.CategoryWithSet, error) {
	query := `
	SELECT 
	categories.id as category_id,
	categories.name as category_name,
	categories.description as category_description,
	sets.id as set_id,
	sets.name as set_name,
	sets.description as set_description 
	FROM categories 
	LEFT JOIN sets ON sets.category_id = categories.id
	WHERE id = $1`
	stmt, err := c.Db.PrepareContext(ctx, query)
	if err != nil {
		c.LoggerService.Error("Failed to prepare query for execution")
		return []models.CategoryWithSet{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	defer stmt.Close()
	var data []models.CategoryWithSet
	sql, err := stmt.QueryContext(ctx, categoryId)
	if err != nil {
		c.LoggerService.Error("Failed to execute query")
		return []models.CategoryWithSet{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	for sql.Next() {
		var categoryWithSet models.CategoryWithSet
		err = sql.Scan(&categoryWithSet.CategoryId, &categoryWithSet.CategoryName, &categoryWithSet.CategoryDescription, &categoryWithSet.SetId, &categoryWithSet.SetName,&categoryWithSet.SetDescription)
		if err != nil{
			c.LoggerService.Error("Failed to scan row")
			return []models.CategoryWithSet{}, models.NewError(500, "Database", "Failed to get data from a database")
		}
		data = append(data, categoryWithSet)
	}
	if err = sql.Err(); err != nil {
		c.LoggerService.Error("Failed to iterate over rows")
		return []models.CategoryWithSet{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return data, nil
}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, categoryId int ) error {
	query :=`
		DELETE FROM categories WHERE id = $1
	`
	stmt, err := c.Db.PrepareContext(ctx, query)
	if err != nil{
		c.LoggerService.Error("Failed to prepare query for execution")
		return models.NewError(500, "Database", "Failed to get data from a database")
	}
	_,err  = stmt.ExecContext(ctx, categoryId)
	if err != nil{
		c.LoggerService.Error("Failed to execute query")
		return models.NewError(500, "Database", "Failed to insert data to a database")
	}
	return nil
}