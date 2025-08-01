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
	sql,err := stmt.QueryContext(ctx, categoryId)
	if err != nil{
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

func(s *SetRepository) GetSetWithElements(ctx context.Context, setId int) (*models.SetWithElements, error) {
	query := `
	SELECT 
		sets.id as id,
		sets.name,
		sets.Description,
		elements.id as elementId,
		elements.English,
		elements.Polish,
		elements.ExampleSentence,
		elements.Synonym
	 FROM sets  
		LEFT JOIN elements ON sets.id = elements.setId 
		WHERE sets.id = $1`
	stmt, err := s.Db.Prepare(query)
	if err != nil{
		s.LoggerService.Error("Failed to prepare query for execution")
		return nil, models.NewError(500, "Database", "Failed to get data from a database")
	}
	sql,err := stmt.QueryContext(ctx, setId)
	if err != nil{
		s.LoggerService.Error("Failed to execute query")
		return nil, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var set models.SetWithElements
	for sql.Next() {
		var element models.Element
		err = sql.Scan(&set.Id, &set.Name, &set.Description,
			&element.Id, &element.English, &element.Polish,
			&element.ExampleSentence, &element.Synonym)
		if err != nil {
			s.LoggerService.Error("Failed to scan row")
			return nil, models.NewError(500, "Database", "Failed to get data from a database")
		}
	}
	if err = sql.Err(); err != nil {
		s.LoggerService.Error("Failed to iterate over rows")
		return nil, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return &set, nil
}
