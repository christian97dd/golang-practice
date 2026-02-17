package main

import (
	"fmt"
)

// Declaración de constantes
// Se recomienda declararlos con mayus los constantes
const Pi float32 = 3.14

const (
	x = 100
	y = 0b1010 // binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	// Declaración de variables
	// var firstName, lastName string
	// var age int

	// var (
	// 	firstName string = "Christian"
	// 	lastName  string = "de Diego"
	// 	age       int    = 28
	// )

	// var firstName, lastName, age = "Christian", "de Diego", 28
	// firstName = "Christian"
	// lastName = "de Diego"
	// age = 28

	firstName, lastName, age := "Christian", "de Diego", 28

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)
}
