package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/middleware"
	"github.com/slodkiadrianek/EI/schema"
)

type SetRoutes struct {
	SetsController *controller.SetsController
}

func NewSetRoutes(setsController *controller.SetsController) *SetRoutes {
	return &SetRoutes{
		SetsController: setsController,
	}
}

func (s *SetRoutes) SetupSetsRouter(router *gin.RouterGroup) {
	sets := router.Group("/sets")
	{
		// sets.GET("/categories/:id", )
		// sets.GET("/:id", )
		sets.POST("/", s.SetsController.CreateSet) // Assuming CreateSet is a method in SetsController
		sets.GET("/", s.SetsController.GetSetsWithElements)
		sets.DELETE("/:setId", middleware.ValidateRequestData[*schema.GetSet]("params"), s.SetsController.DeleteSet)
		// sets.GET("/categories/:id", )
	}
}

