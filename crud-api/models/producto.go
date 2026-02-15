package models

// Producto representa un producto en nuestro sistema
type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre" binding:"required"`
	Precio float64 `json:"precio" binding:"required,gt=0"`
}

// Variable global para almacenar productos en memoria
var Productos []Producto
var SiguienteID int = 1
