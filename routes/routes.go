package routes

import "github.com/gin-gonic/gin"

type SetupRoutes struct {
	SetRoutes       *SetRoutes
	CategoryRoutes *CategoryRoutes
	ElementRoutes *ElementRoutes
}

func (s *SetupRoutes) SetupRouter(router *gin.Engine) {
	routesGroup := router.Group("/api/v1")

	s.SetRoutes.SetupSetsRouter(routesGroup)
	s.CategoryRoutes.SetupCategoriesRouter(routesGroup)
	s.ElementRoutes.SetupElementsRouter(routesGroup)
}

