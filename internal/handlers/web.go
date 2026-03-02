package handlers

import "github.com/gin-gonic/gin"

func RegisterWebRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/tienda", func(c *gin.Context) {
		c.HTML(200, "tienda.html", nil)
	})

	r.GET("/carrito", func(c *gin.Context) {
		c.HTML(200, "carrito.html", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin_login.html", nil)
	})

	r.GET("/admin/productos", AdminOnly(), func(c *gin.Context) {
		c.HTML(200, "admin_productos.html", nil)
	})
}
