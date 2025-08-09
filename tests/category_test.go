package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	router := SetupRouterTest()
	rr := PerformTestRequest(router, "GET", "/api/v1/categories/", nil)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetCategoryById(t *testing.T) {
	router := SetupRouterTest()
	rr := PerformTestRequest(router, "GET", "/api/v1/categories/1", nil)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetCategoryByIdWithSets(t *testing.T) {
	router := SetupRouterTest()
	rr := PerformTestRequest(router, "GET", "/api/v1/categories/1/sets", nil)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateCategory(t *testing.T) {
	router := SetupRouterTest()
	createCategoryData := map[string]string{
		"name":        "TEST",
		"description": "TEST",
	}
	rr := PerformTestRequest(router, "POST", "/api/v1/categories/", createCategoryData)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestDeleteCategory(t *testing.T) {
	router := SetupRouterTest()
	rr := PerformTestRequest(router, "DELETE", "/api/v1/categories/3", nil)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
