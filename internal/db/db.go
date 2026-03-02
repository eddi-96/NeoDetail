package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || name == "" {
		return fmt.Errorf("faltan variables en .env (DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)")
	}

	// 🔹 DSN conexión
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, pass, host, port, name,
	)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error al abrir conexión: %w", err)
	}

	// 🔹 Configuración de pool
	conn.SetMaxOpenConns(15)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(10 * time.Minute)

	// 🔹 Verificar conexión
	if err := conn.Ping(); err != nil {
		return fmt.Errorf("error al hacer ping: %w", err)
	}

	DB = conn
	log.Println("✅ Conectado a MySQL:", host, port, name)

	return nil
}

func SeedSuperUser() {
	query := `
	INSERT INTO usuarios (nombre, email, password_hash, rol)
	SELECT 'UIDE', 'admin@uide.com', '1234', 'ADMIN'
	WHERE NOT EXISTS (
		SELECT 1 FROM usuarios WHERE email = 'admin@uide.com'
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Println("⚠️ Error creando super usuario:", err)
	} else {
		log.Println("👑 Super usuario verificado (UIDE / 1234)")
	}
}
