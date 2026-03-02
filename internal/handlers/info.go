package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var infoOnce sync.Once

func RegisterInfoRoutes(r *gin.Engine) {
	infoOnce.Do(func() {
		r.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "Info OK"})
		})

		r.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	})
}
