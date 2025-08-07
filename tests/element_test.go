package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/config"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/repositories"
	"github.com/slodkiadrianek/EI/routes"
	"github.com/slodkiadrianek/EI/services"
	"github.com/slodkiadrianek/EI/utils"
	"github.com/stretchr/testify/assert"
)

func SetupRouterTest() *gin.Engine {
	router := gin.Default()
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
	routesConfig := routes.SetupRoutes{
		SetRoutes:      routes.NewSetRoutes(SetsController),
		CategoryRoutes: routes.NewCategoryRoutes(CategoryController),
		ElementRoutes:  routes.NewElementRoutes(ElementsController),
	}
	routesConfig.SetupRouter(router)
	return router
}

func TestGetElementsBySetId(t *testing.T) {
	router := SetupRouterTest()
	req, err := http.NewRequest("GET", "/elements/sets/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d got %d", http.StatusOK, status)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
}
