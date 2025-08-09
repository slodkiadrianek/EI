package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// "github.com/slodkiadrianek/EI/tests"
// "github.com/gin-gonic/gin"
// "github.com/slodkiadrianek/EI/config"
// "github.com/slodkiadrianek/EI/controller"
// "github.com/slodkiadrianek/EI/repositories"
// "github.com/slodkiadrianek/EI/routes"
// "github.com/slodkiadrianek/EI/services"
// "github.com/slodkiadrianek/EI/utils"
// "github.com/stretchr/testify/assert"

func TestGetElementsBySetId(t *testing.T) {
	router := SetupRouterTest()
	rr := PerformTestRequest(router, "GET", "/categories/", nil)
	fmt.Print(rr.Code)
	assert.Equal(t, http.StatusOK, rr.Code)
}
