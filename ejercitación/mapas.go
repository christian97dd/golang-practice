package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors["negro"])
	colors["negro"] = "#000000"
	fmt.Println(colors)

	// al obtener la segunda variable, estas obteniendo si existe el valor en el mapa
	// valor, ok := colors["blanco"]
	if valor, ok := colors["verde"]; ok {
		fmt.Println("si existe", valor)
	} else {
		fmt.Println("no existe")
	}
	// fmt.Println(valor, ok)

	// elimino una key
	delete(colors, "negro")
	fmt.Println(colors)

	// itero un map
	for key, value := range colors {
		fmt.Println(key, value)
	}
}
