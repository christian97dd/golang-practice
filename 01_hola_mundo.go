// Este es un comentario de una línea

/*
Este es un comentario
de múltiples líneas
*/

// PACKAGE: Todo archivo Go pertenece a un paquete
// El paquete "main" es especial: indica que este archivo es ejecutable
package main

// IMPORTS: Importamos paquetes que necesitamos
// "fmt" es el paquete de formateo para imprimir en consola
import "fmt"

// FUNCIÓN MAIN: El punto de entrada del programa
// Cuando ejecutas el programa, Go busca esta función y la ejecuta
func main() {
	// Println imprime un mensaje y añade un salto de línea al final
	fmt.Println("¡Hola, Mundo!")

	// Printf permite formatear texto (similar a printf en C)
	// %s es un placeholder para strings (cadenas de texto)
	nombre := "Go" // := declara y asigna una variable (tipo inferido)
	fmt.Printf("Bienvenido a %s\n", nombre)

	// \n es un salto de línea
	// También puedes usar Print (sin salto de línea automático)
	fmt.Print("Este es ")
	fmt.Print("Go\n")
}

/*
PARA EJECUTAR ESTE PROGRAMA:
1. Abre la terminal
2. Navega a esta carpeta: cd go-aprendizaje
3. Ejecuta: go run 01_hola_mundo.go

SALIDA ESPERADA:
¡Hola, Mundo!
Bienvenido a Go
Este es Go
*/
