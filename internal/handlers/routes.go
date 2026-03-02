package handlers

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	RegisterInfoRoutes(r)

	api := r.Group("/api")

	apiAdmin := r.Group("/api/admin")
	apiAdmin.Use(AdminOnly())
	apiAdmin.POST("/productos", ProductoCreate)

	RegisterProductosRoutes(api)
	RegisterServiciosRoutes(api)
	RegisterBeneficiosRoutes(api)
	RegisterCarritoRoutes(api)
	RegisterPedidosRoutes(api)
	RegisterAuthRoutes(api)

}
