package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /api/admin/productos
func ProductoCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ProductoCreate OK",
	})
}
