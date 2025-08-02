package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/controller"
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
		// sets.PUT("/:id", )
		// sets.DELETE("/:id", )
		// sets.GET("/categories/:id", )
	}
	fmt.Print(sets)
}