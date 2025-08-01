package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/utils"
)

type ElementReplository struct{
	LoggerService *utils.Logger
	Db            *sql.DB
}

func NewElementRepository(loggerService *utils.Logger, db *sql.DB) *ElementReplository {
	return &ElementReplository{
		LoggerService: loggerService,
		Db:            db,
	}
}

func (e *ElementReplository) CreateNewElements(ctx context.Context, elements []DTO.Element) error {
	placeholders := []string{}
	values := []string{}
	for i, element := range elements{
		placeholders  = append(placeholders,"($"+string(i*5+1)+", $"+string(i*5+2)+", $"+string(i*5+3)+", $"+string(i*5+4)+", $"+string(i*5+5)+")")
		values= append(values, element.English+", "+element.Polish+", "+element.ExampleSentence+", "+element.Synonym+", "+string(element.SetId))
	}
	query := fmt.Sprintf(`INSERT INTO elements(english, polish, exampleSentence, synonym, setId) VALUES %s`,
		placeholders,
	)
	stmt, err := e.Db.Prepare(query)
	if err != nil {
		e.LoggerService.Error("Failed to prepare query for execution")
		return models.NewError(500, "Database", "Failed to insert data to a database")
	}
	_, err = stmt.ExecContext(ctx, values)
	if err != nil {
		e.LoggerService.Error("Failed to execute query")
		return models.NewError(500, "Database", "Failed to insert data to a database")
	}
	return nil	
}


func (e *ElementReplository) GetElementsFromSet(ctx context.Context, setId int) ([]models.Element, error) {
	query := `
	SELECT 
		id, 
		english,
		polish, 
		exampleSentence, 
		synonym,
	 FROM elements 
	 	WHERE setId = $1`
	stmt, err := e.Db.Prepare(query)
	if err != nil {
		e.LoggerService.Error("Failed to prepare query for execution")
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	sql, err := stmt.QueryContext(ctx, setId)
	if err != nil {
		e.LoggerService.Error("Failed to execute query")
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var elements []models.Element
	for sql.Next() {
		var element models.Element
		err = sql.Scan(&element.Id, &element.English, &element.Polish, &element.ExampleSentence, &element.Synonym, &element.SetId)
		if err != nil {
			e.LoggerService.Error("Failed to scan row")
			return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
		}
		elements = append(elements, element)
	}
	if err = sql.Err(); err != nil {
		e.LoggerService.Error("Failed to iterate over rows")
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return elements, nil
}