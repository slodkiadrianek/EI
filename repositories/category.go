package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/utils"
)

type CategorySet struct{
	LoggerService *utils.Logger
	Db            *sql.DB
}


func (c *CategorySet) CreateCategory(ctx context.Context,category DTO.Category) error{
	query:= "INSERT INTO categories(name,description) VALUES($1,$2)";
	stmt, err:= c.Db.Prepare(query)
	if err !=nil{
		c.LoggerService.Error("Failed to prepare query for execution")
		return models.NewError(500, "Database", "Failed to insert data to a database")
	}
	err = stmt.QueryRowContext(ctx, category).Err();
	if err != nil{
		c.LoggerService.Error("Failed to execute query")
		return models.NewError(500, "Database", "Failed to insert data to a database")
	}
	return nil
}

func (c *CategorySet) GetCategories(ctx context.Context) ([]models.Element, error) {
	query := "SELECT * FROM categories";
	stmt,err := c.Db.Prepare(query)
	if err !=nil{
		c.LoggerService.Error("Failed to prepare query for execution")
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var data []models.Element
	fmt.Println(data);
	err  = stmt.QueryRowContext(ctx).Scan()
	if err != nil{
		c.LoggerService.Error("Failed to execute query")
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return data, nil
}

func (c *CategorySet) GetCategory(ctx context.Context, categoryId int) (models.Element, error) {
	query := "SELECT * FROM categories WHERE id = $1";
	stmt,err := c.Db.Prepare(query)
	if err != nil{
		c.LoggerService.Error("Failed to prepare query for execution")
		return models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var data models.Element
	err = stmt.QueryRowContext(ctx, categoryId).Scan()
	if err != nil{
		c.LoggerService.Error("Failed to execute query")
		return models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return data, nil
}

