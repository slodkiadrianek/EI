package repositories

import (
	"database/sql"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/utils"
	"golang.org/x/net/context"
)

type SetRepository struct {
	LoggerService *utils.Logger
	Db            *sql.DB
}

func NewSetsRepository(loggerService *utils.Logger, db *sql.DB) *SetRepository {
	return &SetRepository{
		LoggerService: loggerService,
		Db:            db,
	}
}

func (s *SetRepository) CreateNewSet(ctx context.Context, set *DTO.Set) (int, error) {
	query := "INSERT INTO sets(name,description,category_id) VALUES($1,$2,$3) RETURNING id"
	stmt, err := s.Db.PrepareContext(ctx, query)
	if err != nil {
		s.LoggerService.Error("Failed to prepare query for execution")
		return 0, models.NewError(500, "Database", err.Error())
	}
	defer stmt.Close()
	var setId int
	err = stmt.QueryRowContext(ctx, set.Name, set.Description, set.CategoryId).Scan(&setId)
	if err != nil {
		s.LoggerService.Error("Failed to execute query")
		return 0, models.NewError(500, "Database", err.Error())
	}
	s.LoggerService.Info("Set created successfully")
	return setId, nil
}

func (s *SetRepository) GetSetsFromCategory(ctx context.Context, categoryId int) ([]models.Set, error) {
	query := "SELECT * FROM sets WHERE category_id = $1"
	stmt, err := s.Db.Prepare(query)
	if err != nil {
		s.LoggerService.Error("Failed to prepare query for execution")
		return []models.Set{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	sql, err := stmt.QueryContext(ctx, categoryId)
	if err != nil {
		s.LoggerService.Error("Failed to execute query")
		return []models.Set{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var sets []models.Set
	for sql.Next() {
		var set models.Set
		err = sql.Scan(&set.Id, &set.Name, &set.Description, &set.CategoryId)
		if err != nil {
			s.LoggerService.Error("Failed to scan row")
			return []models.Set{}, models.NewError(500, "Database", "Failed to get data from a database")
		}
		sets = append(sets, set)
	}
	if err = sql.Err(); err != nil {
		s.LoggerService.Error("Failed to iterate over rows")
		return []models.Set{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return []models.Set{}, nil
}

func (s *SetRepository) GetSets(ctx context.Context,) ([]models.Set, error) {
	query := `
	SELECT 
		sets.id as id,
		sets.name,
		sets.description,
		sets.category_id as categoryId
	 FROM sets  `
	stmt, err := s.Db.PrepareContext(ctx,query)
	if err != nil {
		s.LoggerService.Error("Failed to prepare query for execution")
		return nil, models.NewError(500, "Database", err.Error())
	}
	sql, err := stmt.QueryContext(ctx)
	if err != nil {
		s.LoggerService.Error("Failed to execute query")
		return nil, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var sets []models.Set
	for sql.Next() {
		var set models.Set
		err = sql.Scan(&set.Id, &set.Name, &set.Description, &set.CategoryId)
		if err != nil {
			s.LoggerService.Error("Failed to scan row")
			return nil, models.NewError(500, "Database", "Failed to get data from a database")
		}
		sets = append(sets, set)
	}
	if err = sql.Err(); err != nil {
		s.LoggerService.Error("Failed to iterate over rows")
		return nil, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return sets, nil
}

func (s *SetRepository) DeleteSet(ctx context.Context, setId int) error{
	query :="DELETE FROM sets WHERE id = $1"
	stmt,err := s.Db.PrepareContext(ctx, query)
	if err !=nil{
		s.LoggerService.Error("Failed to prepare query for execution")
		return  models.NewError(500, "Database", "Failed to delete data from a database")
	}
	_,err = stmt.ExecContext(ctx, setId)
	if err != nil{
		s.LoggerService.Error("Failed to execute query")
		return  models.NewError(500, "Database", "Failed to delete data from database")
	}
	return nil
}