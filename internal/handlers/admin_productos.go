package handlers

import (
	"strconv"

	"neodetail/internal/models"
	"neodetail/internal/repository"

	"github.com/gin-gonic/gin"
)

func AdminProductoCrearPost(c *gin.Context) {
	nombre := c.PostForm("nombre")
	precioStr := c.PostForm("precio")
	stockStr := c.PostForm("stock")
	imagen := c.PostForm("imagen_url")
	descripcion := c.PostForm("descripcion")

	precio, _ := strconv.ParseFloat(precioStr, 64)
	stock, _ := strconv.Atoi(stockStr)

	producto := models.Producto{
		Nombre:      nombre,
		Precio:      precio,
		Stock:       stock,
		ImagenURL:   imagen,
		Descripcion: descripcion,
	}

	_, err := repository.CreateProducto(producto)
	if err != nil {
		c.HTML(500, "admin_productos.html", gin.H{"error": "No se pudo crear"})
		return
	}

	c.Redirect(302, "/admin/productos")
}
