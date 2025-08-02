package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)
func ParseCsv(file multipart.File) ([][]string, error) {
    reader := csv.NewReader(file)
    reader.FieldsPerRecord = -1

    var records [][]string
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("error reading csv: %w", err)
        }
        records = append(records, record)
    }
    return records, nil
}

func ExtractValidatedData[T any](source string, c *gin.Context) *T {
	data, _ := c.Get(source)
	typedData, ok := data.(*T)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"category":    "Validation",
				"description": "Proper data does not exist",
			},
		})
		return nil
	}
	return typedData
}