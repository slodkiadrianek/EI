package repositories

import (
	"database/sql"

	"github.com/pelletier/go-toml/query"
	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/config"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/utils"
	"golang.org/x/net/context"
)

type SetRepository struct {
	LoggerService *utils.Logger
	Db            *sql.DB
}

func NewSetsRepository(loggerService *utils.Logger, db *config.Db) *SetsRepository {
	return &SetRepository{
		LoggerService: loggerService,
		Db:            db,
	}
}


func(s *SetRepository) CreateNewSet(ctx context.Context,set *DTO.Set) error {
	query := "INSERT INTO sets(name,description,categoryId) VALUES($1,$2,$3)"
	stmt, err := s.Db.Prepare(query)
	if err != nil{
		s.LoggerService.Error("Failed to prepare query for execution")
		return models.NewError(500, "Database", "Failed to insert data to a database")
	}
	stmt.QueryRowContext(ctx,set)
	return nil
}

func(s *SetRepository)GetSetsFromCategory(ctx context.Context, categoryId int) ([]models.Set ,error) {
	query := "SELECT * FROM sets WHERE categoryId = $1"
	stmt, err := s.Db.Prepare(query)
	if err != nil{
		s.LoggerService.Error("Failed to prepare query for execution")
		return []models.Set{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	err = stmt.QueryRowContext(ctx, categoryId).Scan()
	if err != nil{
		s.LoggerService.Error("Failed to execute query")
		return []models.Set{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return []models.Set{}, nil
}
