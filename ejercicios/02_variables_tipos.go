package main

import "fmt"

func main() {
	fmt.Println("=== VARIABLES Y TIPOS DE DATOS ===\n")

	// ========== DECLARACIÓN DE VARIABLES ==========

	// Forma 1: Declaración explícita con tipo
	var edad int // Declara una variable de tipo entero
	edad = 25     // Asigna un valor
	fmt.Println("Edad:", edad)

	// Forma 2: Declaración e inicialización en una línea
	var nombre string = "Juan"
	fmt.Println("Nombre:", nombre)

	// Forma 3: Inferencia de tipos (Go deduce el tipo automáticamente)
	var ciudad = "Madrid" // Go infiere que es string
	fmt.Println("Ciudad:", ciudad)

	// Forma 4: Declaración corta con := (la más común)
	// Solo funciona DENTRO de funciones
	pais := "España" // Declara y asigna en una línea
	fmt.Println("País:", pais)

	// ========== TIPOS BÁSICOS ==========

	// Enteros (números sin decimales)
	var numeroEntero int = 42
	var numeroEntero8 int8 = 127      // -128 a 127
	var numeroEntero64 int64 = 1000000 // Rango más grande
	fmt.Printf("\nEnteros: %d, %d, %d\n", numeroEntero, numeroEntero8, numeroEntero64)

	// Enteros sin signo (solo positivos)
	var positivo uint = 100 // Solo números >= 0
	fmt.Println("Uint:", positivo)

	// Flotantes (números con decimales)
	var precio float32 = 19.99
	var precioAlto float64 = 999.999 // Más precisión
	fmt.Printf("Flotantes: %.2f, %.3f\n", precio, precioAlto)

	// Booleanos (true o false)
	var esActivo bool = true
	var esMayor bool = false
	fmt.Printf("Booleanos: %t, %t\n", esActivo, esMayor)

	// Strings (cadenas de texto)
	var mensaje string = "Hola Go"
	var multilinea string = `Este es un string
	que puede ocupar
	múltiples líneas` // Usa backticks para multilínea
	fmt.Println("String:", mensaje)
	fmt.Println("Multilínea:", multilinea)

	// ========== CONSTANTES ==========
	// Las constantes NO pueden cambiar su valor
	const PI = 3.14159
	const NOMBRE_APP = "MiApp"
	fmt.Printf("\nConstantes: PI=%.5f, App=%s\n", PI, NOMBRE_APP)

	// ========== MÚLTIPLES VARIABLES ==========
	// Puedes declarar varias variables a la vez
	var x, y, z int = 1, 2, 3
	a, b, c := 10, 20, 30
	fmt.Printf("\nMúltiples: x=%d, y=%d, z=%d\n", x, y, z)
	fmt.Printf("Múltiples cortas: a=%d, b=%d, c=%d\n", a, b, c)

	// ========== VALOR CERO ==========
	// Variables declaradas sin inicializar tienen un "valor cero"
	var enteroVacio int       // 0
	var stringVacio string    // "" (string vacío)
	var boolVacio bool        // false
	fmt.Printf("\nValores cero: int=%d, string='%s', bool=%t\n",
		enteroVacio, stringVacio, boolVacio)

	// ========== CONVERSIÓN DE TIPOS ==========
	// En Go NO hay conversión automática, debes hacerla explícita
	var entero int = 42
	var flotante float64 = float64(entero) // Convierte int a float64
	var otroEntero int = int(flotante)     // Convierte float64 a int
	fmt.Printf("\nConversión: %d -> %.1f -> %d\n", entero, flotante, otroEntero)
}

/*
PUNTOS CLAVE:
1. Go es de tipado estático: cada variable tiene un tipo fijo
2. := es la forma más común de declarar variables (dentro de funciones)
3. Variables sin inicializar tienen un "valor cero" (no undefined/null)
4. Las conversiones de tipo deben ser explícitas
5. Las constantes se declaran con "const" y no cambian

EJECUTAR:
go run 02_variables_tipos.go
*/
