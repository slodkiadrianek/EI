package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/middleware"
	"github.com/slodkiadrianek/EI/schema"
)

type CategoryRoutes struct{
	CategoriesController *controller.CategoryController
}

func NewCategoryRoutes(categoriesController *controller.CategoryController) *CategoryRoutes {
	return &CategoryRoutes{
		CategoriesController: categoriesController,
	}
}


func(c *CategoryRoutes) SetupCategoriesRouter(router *gin.RouterGroup) {
	categories:= router.Group("/categories");
	{
		categories.GET("/",c.CategoriesController.GetCategories) 
		categories.GET("/:categoryId", middleware.ValidateRequestData[*schema.GetCategory]("params"), c.CategoriesController.GetCategory)
		categories.GET("/:categoryId/sets", middleware.ValidateRequestData[*schema.GetCategory]("params"), c.CategoriesController.GetCategoryWithSets)
		categories.POST("/", middleware.ValidateRequestData[*schema.CreateCategory]("body"), c.CategoriesController.CreateCategory)
		categories.DELETE("/:categoryId", middleware.ValidateRequestData[*schema.GetCategory]("params"), c.CategoriesController.DeleteCategory)

	}
}