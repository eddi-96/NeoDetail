package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const EcommerceName = "NeoDetail"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ecommerce": EcommerceName,
			"status":    "API NeoDetail en construcción",
		})
	})

	router.Run(":8080")
}
