package main

import (
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/config"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/middleware"
	"github.com/slodkiadrianek/EI/repositories"
	"github.com/slodkiadrianek/EI/routes"
	"github.com/slodkiadrianek/EI/services"
	"github.com/slodkiadrianek/EI/utils"
)

func main() {
	configEnv := config.SetConfig()
	logger := utils.NewLogger()
	loggerService := logger.CreateLogger()
	Db := config.NewDb(configEnv.DbLink)
	setsRepository := repositories.NewSetsRepository(&loggerService, Db.DbConnection)
	elementsRepository := repositories.NewElementRepository(&loggerService, Db.DbConnection)
	elementService := services.NewElementService(elementsRepository, &loggerService)
	categoryRepository := repositories.NewCategoryRepository(&loggerService, Db.DbConnection)
	categoriesService := services.NewCategoryService(categoryRepository, &loggerService)
	categoryController := controller.NewCategoryController(categoriesService)
	setsService := services.NewSetsService(setsRepository, &loggerService)
	setsController := controller.NewSetsController(setsService, elementService)
	elementsController := controller.NewElementController(elementService)
	router := gin.Default()
	routesConfig := routes.SetupRoutes{
		SetsRoutes:       routes.NewSetsRoutes(setsController),
		CategoriesRoutes: routes.NewcategoriesRoutes(categoryController),
		ElementsRoutes:   routes.NewElementsRoutes(elementsController),
	}

	router.Use(middleware.ErrorMiddleware())
	
	    router.Use(cors.Default())

	routesConfig.SetupRouter(router)
	err := router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	err = router.Run(":3009") // Start the server on port 3009
	if err != nil {
		panic(err)
	}
}
