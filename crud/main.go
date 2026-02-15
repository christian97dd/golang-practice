package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Producto representa un producto en nuestro sistema
type Producto struct {
	ID     int
	Nombre string
	Precio float64
}

// Variable global para almacenar productos (en memoria)
var productos []Producto
var siguienteID int = 1

// CREATE - Crear un nuevo producto
func crearProducto(nombre string, precio float64) {
	nuevoProducto := Producto{
		ID:     siguienteID,
		Nombre: nombre,
		Precio: precio,
	}
	productos = append(productos, nuevoProducto)
	siguienteID++
	fmt.Printf("✓ Producto creado exitosamente con ID: %d\n", nuevoProducto.ID)
}

// READ - Listar todos los productos
func listarProductos() {
	if len(productos) == 0 {
		fmt.Println("No hay productos registrados.")
		return
	}

	fmt.Println("\n=== LISTA DE PRODUCTOS ===")
	fmt.Printf("%-5s %-20s %-10s\n", "ID", "Nombre", "Precio")
	fmt.Println(strings.Repeat("-", 40))
	for _, p := range productos {
		fmt.Printf("%-5d %-20s $%-9.2f\n", p.ID, p.Nombre, p.Precio)
	}
	fmt.Println()
}

// READ - Buscar un producto por ID
func buscarProducto(id int) *Producto {
	for i := range productos {
		if productos[i].ID == id {
			return &productos[i]
		}
	}
	return nil
}

// UPDATE - Actualizar un producto existente
func actualizarProducto(id int, nombre string, precio float64) bool {
	producto := buscarProducto(id)
	if producto == nil {
		fmt.Println("✗ Producto no encontrado.")
		return false
	}

	producto.Nombre = nombre
	producto.Precio = precio
	fmt.Println("✓ Producto actualizado exitosamente.")
	return true
}

// DELETE - Eliminar un producto
func eliminarProducto(id int) bool {
	for i, p := range productos {
		if p.ID == id {
			// Eliminar el producto del slice
			productos = append(productos[:i], productos[i+1:]...)
			fmt.Println("✓ Producto eliminado exitosamente.")
			return true
		}
	}
	fmt.Println("✗ Producto no encontrado.")
	return false
}

// Función auxiliar para leer entrada del usuario
func leerTexto(scanner *bufio.Scanner, mensaje string) string {
	fmt.Print(mensaje)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=================================")
	fmt.Println("   CRUD BÁSICO DE PRODUCTOS")
	fmt.Println("=================================")

	for {
		fmt.Println("\n--- MENÚ ---")
		fmt.Println("1. Crear producto")
		fmt.Println("2. Listar productos")
		fmt.Println("3. Buscar producto por ID")
		fmt.Println("4. Actualizar producto")
		fmt.Println("5. Eliminar producto")
		fmt.Println("6. Salir")

		opcion := leerTexto(scanner, "\nSelecciona una opción: ")

		switch opcion {
		case "1":
			// CREATE
			nombre := leerTexto(scanner, "Nombre del producto: ")
			precioStr := leerTexto(scanner, "Precio: ")
			precio, err := strconv.ParseFloat(precioStr, 64)
			if err != nil {
				fmt.Println("✗ Error: precio inválido")
				continue
			}
			crearProducto(nombre, precio)

		case "2":
			// READ (todos)
			listarProductos()

		case "3":
			// READ (por ID)
			idStr := leerTexto(scanner, "ID del producto: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("✗ Error: ID inválido")
				continue
			}
			producto := buscarProducto(id)
			if producto != nil {
				fmt.Printf("\n--- Producto encontrado ---\n")
				fmt.Printf("ID:     %d\n", producto.ID)
				fmt.Printf("Nombre: %s\n", producto.Nombre)
				fmt.Printf("Precio: $%.2f\n\n", producto.Precio)
			} else {
				fmt.Println("✗ Producto no encontrado.")
			}

		case "4":
			// UPDATE
			idStr := leerTexto(scanner, "ID del producto a actualizar: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("✗ Error: ID inválido")
				continue
			}

			// Primero verificar si existe
			if buscarProducto(id) == nil {
				fmt.Println("✗ Producto no encontrado.")
				continue
			}

			nombre := leerTexto(scanner, "Nuevo nombre: ")
			precioStr := leerTexto(scanner, "Nuevo precio: ")
			precio, err := strconv.ParseFloat(precioStr, 64)
			if err != nil {
				fmt.Println("✗ Error: precio inválido")
				continue
			}
			actualizarProducto(id, nombre, precio)

		case "5":
			// DELETE
			idStr := leerTexto(scanner, "ID del producto a eliminar: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("✗ Error: ID inválido")
				continue
			}
			eliminarProducto(id)

		case "6":
			fmt.Println("\n¡Hasta luego!")
			return

		default:
			fmt.Println("✗ Opción inválida. Por favor intenta de nuevo.")
		}
	}
}
