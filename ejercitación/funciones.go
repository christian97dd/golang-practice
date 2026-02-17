package main

import "fmt"

func main() {
	nombre := hello("juan")
	fmt.Println(nombre)
	sum, mul := calc(4, 5)
	fmt.Println("Operaciones", sum, mul)
}

func hello(name string) string {
	fmt.Println("hola desde la funcion", name)
	return name
}

// funcion que retorna 2 valores y a la vez los creo
// al crearlos e indicar que los va a retornar
// no es necesario hacerlo de manera explícita en el return
func calc(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}
