package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/slodkiadrianek/EI/DTO"
	"github.com/slodkiadrianek/EI/models"
	"github.com/slodkiadrianek/EI/utils"
)

type ElementRepository struct {
	LoggerService *utils.Logger
	Db            *sql.DB
}

func NewElementRepository(loggerService *utils.Logger, db *sql.DB) *ElementRepository {
	return &ElementRepository{
		LoggerService: loggerService,
		Db:            db,
	}
}

func (e *ElementRepository) CreateNewElements(ctx context.Context, elements []DTO.Element) error {
	placeholders := []string{}
	values := []any{}

	for i, element := range elements {
		start := i * 5
		placeholders = append(placeholders,
			"($"+strconv.Itoa(start+1)+", $"+strconv.Itoa(start+2)+", $"+strconv.Itoa(start+3)+", $"+strconv.Itoa(start+4)+", $"+strconv.Itoa(start+5)+")",
		)
		values = append(values,
			element.English,
			element.Polish,
			element.ExampleSentence,
			element.Synonym,
			element.SetId,
		)
	}
	query := fmt.Sprintf(`INSERT INTO elements(english, polish, example_sentence, synonym, set_id) VALUES %s`,
		strings.Join(placeholders, ", "),
	)

	stmt, err := e.Db.PrepareContext(ctx, query)
	if err != nil {
		e.LoggerService.Error("Failed to prepare query for execution", err.Error())
		return models.NewError(500, "Database", "Failed to insert data to the database")
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, values...)
	if err != nil {
		e.LoggerService.Error("Failed to execute query", err.Error())
		return models.NewError(500, "Database", "Failed to insert a data to the database")
	}
	return nil
}

func (e *ElementRepository) GetElementsBySetId(ctx context.Context, setId int) ([]models.Element, error) {
	query := `
	SELECT 
		id, 
		english,
		polish, 
		example_sentence, 
		synonym,
		set_id
	 FROM elements 
	 	WHERE set_id = $1`
	stmt, err := e.Db.Prepare(query)
	if err != nil {
		e.LoggerService.Error("Failed to prepare query for execution", err.Error())
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	defer stmt.Close()
	sql, err := stmt.QueryContext(ctx, setId)
	if err != nil {
		e.LoggerService.Error("Failed to execute query", err.Error())
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var elements []models.Element
	for sql.Next() {
		var element models.Element
		err = sql.Scan(&element.Id, &element.English, &element.Polish, &element.ExampleSentence, &element.Synonym, &element.SetId)
		if err != nil {
			e.LoggerService.Error("Failed to scan row", err.Error())
			return []models.Element{}, models.NewError(500, "Database", err.Error())
		}
		elements = append(elements, element)
	}
	if err = sql.Err(); err != nil {
		e.LoggerService.Error("Failed to iterate over rows", err.Error())
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return elements, nil
}


func (e *ElementRepository) GetElements(ctx context.Context,) ([]models.Element, error) {
	query := `
	SELECT 
		id, 
		english,
		polish, 
		example_sentence, 
		synonym,
		set_id
	 FROM elements `
	stmt, err := e.Db.PrepareContext(ctx,query)
	if err != nil {
		e.LoggerService.Error("Failed to prepare query for execution", err.Error())
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	defer stmt.Close()
	sql, err := stmt.QueryContext(ctx)
	if err != nil {
		e.LoggerService.Error("Failed to execute query", err.Error())
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	var elements []models.Element
	for sql.Next() {
		var element models.Element
		err = sql.Scan(&element.Id, &element.English, &element.Polish, &element.ExampleSentence, &element.Synonym, &element.SetId)
		if err != nil {
			e.LoggerService.Error("Failed to scan row", err.Error())
			return []models.Element{}, models.NewError(500, "Database", err.Error())
		}
		elements = append(elements, element)
	}
	if err = sql.Err(); err != nil {
		e.LoggerService.Error("Failed to iterate over rows", err.Error())
		return []models.Element{}, models.NewError(500, "Database", "Failed to get data from a database")
	}
	return elements, nil
}
