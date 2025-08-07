package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/schema"
	"github.com/slodkiadrianek/EI/services"
	"github.com/slodkiadrianek/EI/utils"
)

type ElementController struct {
	ElementService *services.ElementService
}

func NewElementController(elementService *services.ElementService) *ElementController {
	return &ElementController{
		ElementService: elementService,
	}
}

func (e *ElementController) GetElementsBySetId(c *gin.Context) {
	params := utils.ExtractValidatedData[schema.ElementById]("validatedParams", c)
	elements, err := e.ElementService.GetElementsBySetId(c, params.SetId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{"data": elements})
}
