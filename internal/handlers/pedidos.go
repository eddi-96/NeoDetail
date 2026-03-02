package handlers

import (
	"net/http"

	"neodetail/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterPedidosRoutes(api *gin.RouterGroup) {

	// Obtener pedidos
	api.GET("/pedidos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Pedidos endpoint OK",
		})
	})

	// Finalizar compra
	api.POST("/checkout", func(c *gin.Context) {

		var body struct {
			UsuarioID int `json:"usuario_id"`
		}

		if err := c.ShouldBindJSON(&body); err != nil || body.UsuarioID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "usuario_id inválido",
			})
			return
		}

		pedidoID, err := repository.Checkout(body.UsuarioID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":   true,
			"pedido_id": pedidoID,
			"message":   "Gracias por tu compra",
		})
	})
}
