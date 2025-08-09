package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/config"
	"github.com/slodkiadrianek/EI/controller"
	"github.com/slodkiadrianek/EI/repositories"
	"github.com/slodkiadrianek/EI/routes"
	"github.com/slodkiadrianek/EI/services"
	"github.com/slodkiadrianek/EI/utils"
)

type TestDependecies struct {
	Router             *gin.Engine
	LoggerService      utils.Logger
	Db                 *config.Db
	SetsRepository     repositories.SetRepository
	ElementsRepository repositories.ElementRepository
	CategoryRepository repositories.CategoryRepository
	ElementService     services.ElementService
	CategoriesService  services.CategoryService
	SetService         services.SetService
	CategoryController controller.CategoryController
	SetsController     controller.SetsController
	ElementsController controller.ElementController
}

func NewDependecies() *TestDependecies {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	configEnv := config.SetConfig("../.env.test")
	logger := utils.NewLogger()
	loggerService := logger.CreateLogger()
	fmt.Println(configEnv.DbLink)
	db := config.NewDb(configEnv.DbLink)
	setsRepository := repositories.NewSetRepository(&loggerService, db.DbConnection)
	elementsRepository := repositories.NewElementRepository(&loggerService, db.DbConnection)
	elementService := services.NewElementService(elementsRepository, &loggerService)
	categoryRepository := repositories.NewCategoryRepository(&loggerService, db.DbConnection)
	categoriesService := services.NewCategoryService(categoryRepository, &loggerService)
	categoryController := controller.NewCategoryController(categoriesService)
	setService := services.NewSetService(setsRepository, elementsRepository, &loggerService)
	setsController := controller.NewSetsController(setService, elementService)
	elementsController := controller.NewElementController(elementService)
	routesConfig := routes.SetupRoutes{
		SetRoutes:      routes.NewSetRoutes(setsController),
		CategoryRoutes: routes.NewCategoryRoutes(categoryController),
		ElementRoutes:  routes.NewElementRoutes(elementsController),
	}
	routesConfig.SetupRouter(router)
	return &TestDependecies{
		Router:             router,
		LoggerService:      loggerService,
		Db:                 db,
		SetsRepository:     *setsRepository,
		ElementsRepository: *elementsRepository,
		CategoryRepository: *categoryRepository,
		ElementService:     *elementService,
		CategoriesService:  *categoriesService,
		SetService:         *setService,
		CategoryController: *categoryController,
		SetsController:     *setsController,
		ElementsController: *elementsController,
	}
}

func SetupRouterTest() *gin.Engine {
	return NewDependecies().Router
}

func CreateTestRequest(method, url string, body any) (*http.Request, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(bodyBytes)
	return http.NewRequest(method, url, reader)
}

func PerformTestRequest(router *gin.Engine, method, url string, body any) *httptest.ResponseRecorder {
	req, _ := CreateTestRequest(method, url, body)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
