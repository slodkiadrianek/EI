package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/middleware"
	"github.com/slodkiadrianek/EI/schema"
)

type ElementsRoutes struct{
	ElementsController *controller.ElementController
}

func NewElementsRoutes(ElementsController *controller.ElementController) *ElementsRoutes {
	return &ElementsRoutes{
		ElementsController: ElementsController,
	}
}


func(s *ElementsRoutes) SetupElementsRouter(router *gin.RouterGroup) {
	elements:= router.Group("/elements");
	{
		// Elements.GET("/categories/:id", )
		// Elements.GET("/:id", )
		elements.GET("/sets/:setId",middleware.ValidateRequestData[*schema.ElementsById]("params"),s.ElementsController.GetElementsBySetId) // Assuming CreateSet is a method in ElementsController
		// Elements.PUT("/:id", )
		// Elements.DELETE("/:id", )
		// Elements.GET("/categories/:id", )
	}
	fmt.Print(elements)
}