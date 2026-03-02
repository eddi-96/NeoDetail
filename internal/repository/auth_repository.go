package repository

import (
	"database/sql"
	"neodetail/internal/db"
)

func ValidateAdmin(usuario, pass string) (bool, error) {
	// tú guardaste UIDE / 1234, y rol ADMIN
	var rol string
	err := db.DB.QueryRow(`
		SELECT rol FROM usuarios
		WHERE nombre = ? AND password_hash = ? LIMIT 1
	`, usuario, pass).Scan(&rol)

	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return rol == "ADMIN", nil
}
