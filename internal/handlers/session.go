package handlers

import "github.com/gin-gonic/gin"

func LoadSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Cookie("role")
		if role != "" {
			c.Set("role", role)
		}
		c.Next()
	}
}
