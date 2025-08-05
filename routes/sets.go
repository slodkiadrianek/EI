package routes

import (

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/middleware"
	"github.com/slodkiadrianek/EI/schema"
)

type SetsRoutes struct{
	SetsController *controller.SetsController
}

func NewSetsRoutes(setsController *controller.SetsController) *SetsRoutes {
	return &SetsRoutes{
		SetsController: setsController,
	}
}


func(s *SetsRoutes) SetupSetsRouter(router *gin.RouterGroup) {
	sets:= router.Group("/sets");
	{
		// sets.GET("/categories/:id", )
		// sets.GET("/:id", )
		sets.POST("/",s.SetsController.CreateSet) // Assuming CreateSet is a method in SetsController
		sets.GET("/", s.SetsController.GetSetsWithElements)
		sets.DELETE("/:id",middleware.ValidateRequestData[*schema.GetSet]("params"), s.SetsController.DeleteSet)
		// sets.GET("/categories/:id", )
	}
}