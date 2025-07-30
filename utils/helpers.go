package utils

import (
	"encoding/csv"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseCsv(file multipart.File) ([][]string, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
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