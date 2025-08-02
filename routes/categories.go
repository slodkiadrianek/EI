package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/middleware"
	"github.com/slodkiadrianek/EI/schema"
)

type CategoriesRoutes struct{
	CategoriesController *controller.CategoryController
}

func NewcategoriesRoutes(categoriesController *controller.CategoryController) *CategoriesRoutes {
	return &CategoriesRoutes{
		CategoriesController: categoriesController,
	}
}


func(c *CategoriesRoutes) SetupCategoriesRouter(router *gin.RouterGroup) {
	categories:= router.Group("/categories");
	{
		categories.GET("/",c.CategoriesController.GetCategories) 
		categories.GET("/:categoryId", middleware.ValidateRequestData[*schema.GetCategory]("params"), c.CategoriesController.GetCategory)
		categories.POST("/", middleware.ValidateRequestData[*schema.CreateCategory]("body"), c.CategoriesController.CreateCategory)
	}
}