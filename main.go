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
	ginSwagger "github.com/swaggo/gin-swagger"
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
	categoryController := controller.NewCategoryController(categoriesService)
	SetService := services.NewSetService(setsRepository, elementsRepository,&loggerService)
	setsController := controller.NewSetsController(SetService, elementService)
	elementsController := controller.NewElementController(elementService)
	router := gin.Default()
	routesConfig := routes.SetupRoutes{
		SetRoutes:       routes.NewSetRoutes(setsController),
		CategoryRoutes: routes.NewCategoryRoutes(categoryController),
		ElementRoutes:   routes.NewElementRoutes(elementsController),
	}

	router.Use(middleware.ErrorMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.Default())
	routesConfig.SetupRouter(router)
	err := router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	err = router.Run(":"+configEnv.Port) 
	if err != nil {
		panic(err)
	}
}
