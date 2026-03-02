package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != "admin" {
			c.Redirect(http.StatusFound, "/admin")
			c.Abort()
			return
		}
		c.Next()
	}
}
