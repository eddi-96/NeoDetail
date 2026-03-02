package handlers

import (
	"neodetail/internal/models"
	"neodetail/internal/repository"

	"github.com/gin-gonic/gin"
)

func ProductoCreate(c *gin.Context) {
	var producto models.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	_, err := repository.CreateProducto(producto)
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el producto"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Producto creado correctamente",
	})
}
func RegisterProductosRoutes(r *gin.RouterGroup) {
	r.GET("/productos", ProductoList)
	r.POST("/productos", ProductoCreate)
}

func ProductoList(c *gin.Context) {
	productos, err := repository.GetProductos()
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudieron obtener productos"})
		return
	}

	c.JSON(200, productos)
}
