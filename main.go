package main

import (
	"log"

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
	setsRepository := repositories.NewSetsRepository(&loggerService, Db)
	setsService := services.NewSetsService(setsRepository, &loggerService)
	setsController := controller.NewSetsController(setsService)
	router := gin.Default()
	routesConfig := routes.SetupRoutes{
		SetsRoutes: routes.NewSetsRoutes(setsController),
		// CategoriesRoutes: routes.NewCategoriesRoutes(),
	}

	router.Use(middleware.ErrorMiddleware())

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
