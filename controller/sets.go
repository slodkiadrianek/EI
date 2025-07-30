package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/services"
)

type SetsController struct {
	SetsService *services.SetsService
}

func NewSetsController(setsService *services.SetsService) *SetsController {
	return &SetsController{
		SetsService: setsService,
	}
}

func GetSets(c *gin.Context) {
	// Logic to get all sets
	c.JSON(200, gin.H{"message": "Get all sets"})
}

func GetSetByID(c *gin.Context) {
	// Logic to get a set by ID
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get set by ID", "id": id})
}

func (s *SetsController) CreateSet(c *gin.Context) {
	file, _, err := c.Request.FormFile("csv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error: " + err.Error()})
	}
	defer file.Close()
	c.JSON(201, gin.H{})
}
