package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/slodkiadrianek/EI/schema"
	"github.com/slodkiadrianek/EI/services"
	"github.com/slodkiadrianek/EI/utils"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	data := utils.ExtractValidatedData[schema.CreateCategory]("validatedData", c)
	err := cc.CategoryService.CreateCategory(c, data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, gin.H{})
}

func (cc *CategoryController) GetCategories(c *gin.Context) {
	categories, err := cc.CategoryService.GetCategories(c)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{
		"categories": categories,
	})
}


func (cc *CategoryController) GetCategoryWithSets(c *gin.Context) {
	categoryParams := utils.ExtractValidatedData[schema.GetCategory]("validatedParams", c)
	categoriesWithSets,err :=cc.CategoryService.GetCategoryWithSets(c, categoryParams.CategoryId)
	if err != nil{
		c.Error(err)
		return 
	}
	c.JSON(200, gin.H{"data":categoriesWithSets})
}


func (cc *CategoryController) DeleteCategory(c *gin.Context){
	categoryParams := utils.ExtractValidatedData[schema.GetCategory]("validatedParams", c)
	err := cc.CategoryService.DeleteCategory(c, categoryParams.CategoryId)
	if err !=nil{
		c.Error(err)
		return
	}
	c.JSON(204,gin.H{})
}
