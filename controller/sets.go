package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/schema"
	"github.com/slodkiadrianek/EI/services"
)

type SetsController struct {
	SetsService     *services.SetsService
	ElementsService *services.ElementService
}

func NewSetsController(setsService *services.SetsService, elementsService *services.ElementService) *SetsController {
	return &SetsController{
		SetsService:     setsService,
		ElementsService: elementsService,
	}
}

func GetSets(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get all sets"})
}

func GetSetByID(c *gin.Context) {
	// Logic to get a set by ID
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get set by ID", "id": id})
}

func (s *SetsController) CreateSet(c *gin.Context) {
	name := c.Request.FormValue("name")
	description := c.Request.FormValue("description")
	categoryId := c.Request.FormValue("categoryId")
	var data schema.CreateSet
	data.Name = name
	data.Description = description
	val, err := strconv.Atoi(categoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	data.CategoryId = val
	file, _, err := c.Request.FormFile("csv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error: " + err.Error()})
	}
	defer file.Close()
	setId, err := s.SetsService.CreateSet(c, data.CategoryId, &data)
	if err != nil {
		c.Error(err)
		return
	}
	err = s.ElementsService.CreateElements(c, file, "name", setId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, gin.H{})
}
