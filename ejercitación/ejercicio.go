package main

/*
Ejercicio: Calcule e imprima el área y el perímetro del triángulo
Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.



El programa debe:



Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.

Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.

Calcular el área del triángulo usando la fórmula base x altura / 2.

Calcular el perímetro del triángulo sumando los lados.

Imprimir el área y el perímetro del triángulo con dos decimales de precisión.

El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro. También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.

Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.



Ejemplo de entrada y salida:

Ingrese lado 1: 3.5
Ingrese lado 2: 4.2

Área: 7.35
Perímetro: 12.20

*/

import (
	"fmt"
	"math"
)

func main() {
	var cateto1 float64
	var cateto2 float64
	const precision int = 2

	fmt.Print("ingrese longitud cateto 1 del triangulo rectangulo\n")
	fmt.Scanln(&cateto1)

	fmt.Print("ingrese longitud cateto 2 del triangulo rectangulo\n")
	fmt.Scanln(&cateto2)

	sumaLadosCuadrados := math.Pow(cateto1, 2) + math.Pow(cateto2, 2)

	hipotenusa := math.Sqrt(sumaLadosCuadrados)

	areaTriangulo := (cateto1 * cateto2) / 2
	perimetroTriangulo := cateto1 + cateto2 + hipotenusa

	fmt.Println("el lado 3 mide:", hipotenusa)
	fmt.Println("el area del triangulo es:", areaTriangulo)
	fmt.Println("el perímetro del triangulo es:", perimetroTriangulo)

	fmt.Printf("el area del triangulo con 2 digitos es: %.*f\n", precision, areaTriangulo)
	fmt.Printf("el perímetro del triangulo con 2 digitos es: %.*f\n", precision, perimetroTriangulo)
}
