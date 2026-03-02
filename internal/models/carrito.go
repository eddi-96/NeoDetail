package models

type Carrito struct {
	ID        int           `json:"id"`
	UsuarioID int           `json:"usuario_id"`
	Estado    string        `json:"estado"`
	Items     []CarritoItem `json:"items"`
}

type CarritoItem struct {
	ID             int     `json:"id"`
	CarritoID      int     `json:"carrito_id"`
	ProductoID     int     `json:"producto_id"`
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`

	// ✅ CAMPOS QUE TE FALTAN (para mostrar en la UI)
	ProductoNombre string `json:"producto_nombre"`
	ImagenURL      string `json:"imagen_url"`
}
