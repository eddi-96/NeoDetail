package repository

import (
	"neodetail/internal/db"
	"neodetail/internal/models"
)

func GetProductos() ([]models.Producto, error) {
	rows, err := db.DB.Query(`
		SELECT id, nombre, descripcion, precio, stock, imagen_url
		FROM productos WHERE activo = TRUE
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []models.Producto
	for rows.Next() {
		var p models.Producto
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.Stock, &p.ImagenURL); err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}

func CreateProducto(p models.Producto) (int64, error) {
	res, err := db.DB.Exec(`
		INSERT INTO productos (nombre, descripcion, precio, stock, imagen_url, activo)
		VALUES (?, ?, ?, ?, ?, TRUE)
	`, p.Nombre, p.Descripcion, p.Precio, p.Stock, p.ImagenURL)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
