package handlers

import "github.com/gin-gonic/gin"

func RegisterBeneficiosRoutes(api *gin.RouterGroup) {
	api.GET("/beneficios", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "Beneficios (siguiente: SQL)", "data": []any{}})
	})
}
