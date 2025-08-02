package routes

import "github.com/gin-gonic/gin"

type SetupRoutes struct {
	SetsRoutes       *SetsRoutes
	CategoriesRoutes *CategoriesRoutes
}

func (s *SetupRoutes) SetupRouter(router *gin.Engine) {
	routesGroup := router.Group("/api/v1")

	s.SetsRoutes.SetupSetsRouter(routesGroup)
	s.CategoriesRoutes.SetupCategoriesRouter(routesGroup)
}

