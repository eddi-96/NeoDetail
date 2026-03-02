package handlers

import (
	"net/http"
	"strings"

	"neodetail/internal/repository"

	"github.com/gin-gonic/gin"
)

func AdminLoginPost(c *gin.Context) {
	usuario := strings.TrimSpace(c.PostForm("usuario"))
	clave := strings.TrimSpace(c.PostForm("password"))

	if usuario == "" || clave == "" {
		c.HTML(http.StatusBadRequest, "admin_login.html", gin.H{"error": "Completa usuario y contraseña"})
		return
	}

	ok, err := repository.ValidateAdmin(usuario, clave)
	if err != nil || !ok {
		c.HTML(http.StatusUnauthorized, "admin_login.html", gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// cookie simple (para tarea)
	c.SetCookie("admin", "1", 3600, "/", "", false, true)
	c.Redirect(http.StatusFound, "/admin/productos")
}
