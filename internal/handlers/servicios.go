package handlers

import "github.com/gin-gonic/gin"

func RegisterServiciosRoutes(api *gin.RouterGroup) {
	api.GET("/servicios", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "Servicios (siguiente: SQL)", "data": []any{}})
	})
}
