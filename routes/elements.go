package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/middleware"
	"github.com/slodkiadrianek/EI/schema"
)

type ElementRoutes struct{
	ElementsController *controller.ElementController
}

func NewElementRoutes(ElementsController *controller.ElementController) *ElementRoutes {
	return &ElementRoutes{
		ElementsController: ElementsController,
	}
}


func(s *ElementRoutes) SetupElementsRouter(router *gin.RouterGroup) {
	elements:= router.Group("/elements");
	{
		// Elements.GET("/categories/:id", )
		// Elements.GET("/:id", )
		elements.GET("/sets/:setId",middleware.ValidateRequestData[*schema.ElementById]("params"),s.ElementsController.GetElementsBySetId) // Assuming CreateSet is a method in ElementsController
		// Elements.PUT("/:id", )
		// Elements.DELETE("/:id", )
		// Elements.GET("/categories/:id", )
	}
}