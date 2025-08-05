package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/schema"
	"github.com/slodkiadrianek/EI/services"
	"github.com/slodkiadrianek/EI/utils"
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

func (s *SetsController) GetSetsWithElements(c *gin.Context){
	setsWithElements, err :=s.SetsService.GetSetsWithElements(c)
	if err != nil{
		c.Error(err)
		return 
	}
	c.JSON(200, gin.H{"data": setsWithElements})
}

func(s *SetsController) DeleteSet(c *gin.Context) {
	params := utils.ExtractValidatedData[schema.GetSet]("validatedParams", c)
	err := s.SetsService.DeleteSet(c, params.SetId)
	if err != nil{
		c.Error(err)
		return 
	}
	c.JSON(204, gin.H{})
}