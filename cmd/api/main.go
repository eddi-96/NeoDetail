package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"neodetail/internal/db"
	"neodetail/internal/handlers"
)

func main() {

	// 🔹 Cargar variables .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No se encontró archivo .env, usando variables del sistema")
	}

	// 🔹 Conectar BD
	if err := db.Connect(); err != nil {
		log.Fatal("❌ Error conectando a MySQL: ", err)
	}

	// ✅ Crear/Verificar SUPER USUARIO (UIDE / 1234)
	db.SeedSuperUser()

	// 🔹 Crear router
	r := gin.Default()
	r.Use(handlers.LoadSession())

	// ===============================
	// ✅ WEB (templates + estáticos)
	// ===============================
	r.LoadHTMLGlob("web/templates/*.html")
	r.Static("/static", "./web/static")

	// ===============================
	// ✅ Rutas (UNA sola vez)
	// ===============================
	handlers.RegisterRoutes(r)    // API
	handlers.RegisterWebRoutes(r) // WEB

	// 🔹 Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Servidor corriendo en http://localhost:" + port)
	_ = r.Run(":" + port)
}
