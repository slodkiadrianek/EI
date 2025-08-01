package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/slodkiadrianek/EI/models"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				var customErr *models.ErrorResponse
				if errors.As(e.Err, &customErr) {
					c.JSON(customErr.Code, gin.H{
						"error": gin.H{
							"category":    customErr.Category,
							"description": customErr.Description,
						},
					})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error": gin.H{
						"category":    "Internal Server Error",
						"description": e.Error(),
					},
				})
			}
		}
	}
}