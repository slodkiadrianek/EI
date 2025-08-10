package main

import (
	"fmt"
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

	setsRepository := repositories.NewSetRepository(&loggerService, Db.DbConnection)
	elementsRepository := repositories.NewElementRepository(&loggerService, Db.DbConnection)
	elementService := services.NewElementService(elementsRepository, &loggerService)
	categoryRepository := repositories.NewCategoryRepository(&loggerService, Db.DbConnection)
	categoriesService := services.NewCategoryService(categoryRepository, &loggerService)
	CategoryController := controller.NewCategoryController(categoriesService)
	SetService := services.NewSetService(setsRepository, elementsRepository, &loggerService)
	SetsController := controller.NewSetsController(SetService, elementService)
	ElementsController := controller.NewElementController(elementService)
	router := gin.Default()
	routesConfig := routes.SetupRoutes{
		SetRoutes:      routes.NewSetRoutes(SetsController),
		CategoryRoutes: routes.NewCategoryRoutes(CategoryController),
		ElementRoutes:  routes.NewElementRoutes(ElementsController),
	}

	router.Use(middleware.ErrorMiddleware())
	router.Use(cors.Default())
	routesConfig.SetupRouter(router)
	fmt.Println(configEnv)
	fmt.Println("Server is running on port:", configEnv.Port)
	fmt.Println(configEnv.Port)
	err := router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	err = router.Run(":" + configEnv.Port)
	if err != nil {
		panic(err)
	}
}
