package repository

import (
	"database/sql"
	"errors"
)

type Carrito struct {
	ID        int
	UsuarioID int
	Estado    string
}

type CarritoItem struct {
	ID         int
	CarritoID  int
	ProductoID int
	Cantidad   int
}

func GetCarritoByUsuario(usuarioID int) (*Carrito, error) {
	// Busca carrito ABIERTO del usuario (ejemplo)
	const q = `
		SELECT TOP 1 id, usuario_id, estado
		FROM carritos
		WHERE usuario_id = ? AND estado = 'ABIERTO'
		ORDER BY id DESC;
	`

	var c Carrito
	err := DB.QueryRow(q, usuarioID).Scan(&c.ID, &c.UsuarioID, &c.Estado)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &c, nil
}

func GetOrCreateCarritoAbierto(usuarioID int) (*Carrito, error) {
	// 1) intenta obtener
	c, err := GetCarritoByUsuario(usuarioID)
	if err == nil {
		return c, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	// 2) crea si no existe
	const insertQ = `
		INSERT INTO carritos (usuario_id, estado)
		OUTPUT INSERTED.id, INSERTED.usuario_id, INSERTED.estado
		VALUES (?, 'ABIERTO');
	`

	var nuevo Carrito
	err = DB.QueryRow(insertQ, usuarioID).Scan(&nuevo.ID, &nuevo.UsuarioID, &nuevo.Estado)
	if err != nil {
		return nil, err
	}
	return &nuevo, nil
}

func AddItemToCarrito(carritoID, productoID, cantidad int) error {
	// Lógica típica: si existe item, sumar; si no existe, insertar

	// 1) ver si existe
	const findQ = `
		SELECT TOP 1 id, cantidad
		FROM carrito_items
		WHERE carrito_id = ? AND producto_id = ?;
	`

	var itemID int
	var cantActual int
	err := DB.QueryRow(findQ, carritoID, productoID).Scan(&itemID, &cantActual)

	if err == nil {
		// actualiza
		const updQ = `
			UPDATE carrito_items
			SET cantidad = ?
			WHERE id = ?;
		`
		_, err = DB.Exec(updQ, cantActual+cantidad, itemID)
		return err
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// insertar nuevo
	const insQ = `
		INSERT INTO carrito_items (carrito_id, producto_id, cantidad)
		VALUES (?, ?, ?);
	`
	_, err = DB.Exec(insQ, carritoID, productoID, cantidad)
	return err
}
