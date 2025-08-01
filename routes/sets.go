package routes

import "github.com/gin-gonic/gin"

type SetsRoutes struct{
	SetsController interface{}
}

func NewSetsRoutes(setsController interface{}) *SetsRoutes {
	return &SetsRoutes{
		SetsController: setsController,
	}
}


func(s *SetsRoutes) SetupSetsRouter(router *gin.RouterGroup) {
	sets:= router.Group("/sets");
	{
		sets.GET("/categories/:id", )
		sets.GET("/:id", )
		sets.POST("/", )
		sets.PUT("/:id", )
		sets.DELETE("/:id", )
		sets.GET("/categories/:id", )
	}
}