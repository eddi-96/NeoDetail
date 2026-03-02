package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// Checkout: toma el carrito ABIERTO del usuario y genera un pedido.
// Devuelve el pedidoID.
func Checkout(usuarioID int) (int, error) {
	if DB == nil {
		return 0, errors.New("DB no inicializada en repository")
	}

	// 1) Obtener carrito abierto
	carrito, err := GetCarritoByUsuario(usuarioID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("no hay carrito abierto para usuario %d", usuarioID)
		}
		return 0, err
	}

	// 2) Verificar que tenga items
	var totalItems int
	err = DB.QueryRow(`
		SELECT COUNT(*)
		FROM carrito_items
		WHERE carrito_id = ?;
	`, carrito.ID).Scan(&totalItems)

	if err != nil {
		return 0, err
	}
	if totalItems == 0 {
		return 0, errors.New("el carrito está vacío")
	}

	// 3) Transacción (para que sea seguro)
	tx, err := DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = tx.Rollback() // si ya se hizo commit no hace nada
	}()

	// 4) Crear pedido
	var pedidoID int
	err = tx.QueryRow(`
		INSERT INTO pedidos (usuario_id, fecha, estado)
		OUTPUT INSERTED.id
		VALUES (?, ?, 'CREADO');
	`, usuarioID, time.Now()).Scan(&pedidoID)
	if err != nil {
		return 0, err
	}

	// 5) Copiar items del carrito a detalle_pedido
	_, err = tx.Exec(`
		INSERT INTO detalle_pedido (pedido_id, producto_id, cantidad)
		SELECT ?, producto_id, cantidad
		FROM carrito_items
		WHERE carrito_id = ?;
	`, pedidoID, carrito.ID)
	if err != nil {
		return 0, err
	}

	// 6) Cerrar carrito (cambiar estado)
	_, err = tx.Exec(`
		UPDATE carritos
		SET estado = 'CERRADO'
		WHERE id = ?;
	`, carrito.ID)
	if err != nil {
		return 0, err
	}

	// 7) Commit
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return pedidoID, nil
}
