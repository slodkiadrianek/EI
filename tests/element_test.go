package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetElementsBySetId(t *testing.T) {
	router := SetupRouterTest()
	rr := PerformTestRequest(router, "GET", "/api/v1/elements/sets/1", nil)
	assert.Equal(t, http.StatusOK, rr.Code)
}
