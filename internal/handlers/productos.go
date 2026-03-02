package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterProductosRoutes(api *gin.RouterGroup) {
	// GET /api/productos
	api.GET("/productos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "RegisterProductosRoutes OK",
		})
	})
}
