package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSets(t *testing.T) {
	router := SetupRouterTest()
	rr := PerformTestRequest(router, "GET", "/api/v1/sets/", nil)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetSetById(t *testing.T) {
	router := SetupRouterTest()
	rr := PerformTestRequest(router, "DELETE", "/api/v1/sets/1", nil)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
