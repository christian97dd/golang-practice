package main

import "fmt"

type Persona2 struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var persona Persona2
	persona.nombre = "Christian"
	persona.edad = 28
	persona.correo = "exequieldediego@gmail.com"
	fmt.Println(persona)

	persona2 := Persona2{"cdd", 28, "fafafa"}
	fmt.Println(persona2)
}
