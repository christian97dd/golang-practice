package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) sayHi() {
	fmt.Println("hola, mi nombre es", p.nombre)
}

func main() {
	var x int = 10
	fmt.Println(x)
	editar(&x)
	var p *int = &x

	fmt.Println(x, p)

	// instancio 2 personas, cada uno va a tener su propio metodo sayHi
	persona := Persona{"Christian", 28, "cdd@gmail.com"}
	persona.sayHi()

	persona2 := Persona{"aaa", 28, "cdd@gmail.com"}
	persona2.sayHi()
}

func editar(x *int) {
	*x = 20
}
