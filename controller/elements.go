package controller

import (
	"github.com/slodkiadrianek/EI/services"
)

type ElementController struct {
	ElementService *services.ElementService
}

func NewElementController(elementService *services.ElementService) *ElementController {
	return &ElementController{
		ElementService: elementService,
	}
}
