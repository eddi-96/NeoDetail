package handlers

import (
	"net/http"
	"strconv"

	"neodetail/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterCarritoRoutes(api *gin.RouterGroup) {

	// ✅ Obtener carrito por usuario
	api.GET("/carrito", func(c *gin.Context) {

		userIDStr := c.Query("usuario_id")
		if userIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "usuario_id es requerido",
			})
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil || userID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "usuario_id inválido",
			})
			return
		}

		carrito, err := repository.GetCarritoByUsuario(userID)

		// ✅ Si NO existe carrito abierto, devolvemos vacío (200 OK)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data": gin.H{
					"id":         0,
					"usuario_id": userID,
					"estado":     "ABIERTO",
					"items":      []any{},
				},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    carrito,
		})
	})

	// ✅ Agregar producto al carrito
	api.POST("/carrito/add", func(c *gin.Context) {

		var body struct {
			UsuarioID  int `json:"usuario_id"`
			ProductoID int `json:"producto_id"`
			Cantidad   int `json:"cantidad"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "JSON inválido"})
			return
		}

		if body.UsuarioID <= 0 || body.ProductoID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "usuario_id y producto_id son obligatorios"})
			return
		}

		if body.Cantidad <= 0 {
			body.Cantidad = 1
		}

		// 1) Obtener o crear carrito ABIERTO
		carrito, err := repository.GetOrCreateCarritoAbierto(body.UsuarioID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}

		// 2) Insertar item (o sumar si ya existe)
		if err := repository.AddItemToCarrito(carrito.ID, body.ProductoID, body.Cantidad); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Producto agregado al carrito",
			"carrito_id": carrito.ID,
		})
	})
}
