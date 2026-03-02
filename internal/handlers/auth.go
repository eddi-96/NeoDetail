package handlers

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(api *gin.RouterGroup) {
	api.POST("/auth/register", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "Register (mock)"})
	})
	api.POST("/auth/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "message": "Login (mock)"})
	})
}
