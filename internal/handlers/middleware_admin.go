package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		val, err := c.Cookie("admin")
		if err != nil || val != "1" {
			c.Redirect(http.StatusFound, "/admin")
			c.Abort()
			return
		}
		c.Next()
	}
}
