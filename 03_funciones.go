package main

import "fmt"

// ========== FUNCIÓN SIMPLE ==========
// Declaración: func nombre(parámetros) tipoRetorno { cuerpo }
func saludar() {
	fmt.Println("¡Hola desde una función!")
}

// ========== FUNCIÓN CON PARÁMETROS ==========
// Los parámetros llevan nombre y tipo
func saludarPersona(nombre string) {
	fmt.Printf("Hola, %s\n", nombre)
}

// ========== FUNCIÓN CON RETORNO ==========
// El tipo de retorno va después de los parámetros
func sumar(a int, b int) int {
	return a + b
}

// Si varios parámetros son del mismo tipo, puedes agruparlos
func multiplicar(a, b int) int {
	return a * b
}

// ========== MÚLTIPLES VALORES DE RETORNO ==========
// Go permite retornar varios valores (muy útil para errores)
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// Retornamos 0 y un error
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	// Retornamos el resultado y nil (sin error)
	return a / b, nil
}

// ========== VALORES DE RETORNO NOMBRADOS ==========
// Puedes nombrar los valores de retorno
func rectangulo(base, altura float64) (area, perimetro float64) {
	// area y perimetro ya están declarados
	area = base * altura
	perimetro = 2 * (base + altura)
	return // Retorna automáticamente area y perimetro
}

// ========== FUNCIONES VARIÁDICAS ==========
// Aceptan un número variable de argumentos (como ...rest en JS)
func sumarTodos(numeros ...int) int {
	total := 0
	// numeros es un slice (array) de enteros
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// ========== FUNCIONES COMO VALORES ==========
// Las funciones son "first-class citizens" (como en JavaScript)
func aplicarOperacion(a, b int, operacion func(int, int) int) int {
	return operacion(a, b)
}

// ========== FUNCIONES ANÓNIMAS ==========
func ejemploFuncionAnonima() {
	// Puedes declarar funciones sin nombre
	saludar := func(nombre string) {
		fmt.Printf("Hola, %s (desde función anónima)\n", nombre)
	}
	saludar("María")
}

// ========== CLOSURES ==========
// Las funciones pueden capturar variables de su entorno
func contador() func() int {
	count := 0 // Variable capturada
	return func() int {
		count++ // Modifica la variable capturada
		return count
	}
}

func main() {
	fmt.Println("=== FUNCIONES EN GO ===\n")

	// Llamar a función simple
	saludar()

	// Función con parámetros
	saludarPersona("Carlos")

	// Función con retorno
	resultado := sumar(5, 3)
	fmt.Printf("5 + 3 = %d\n", resultado)

	// Múltiples retornos
	division, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", division)
	}

	// Intentar dividir por cero
	_, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Error capturado:", err)
	}

	// Valores de retorno nombrados
	area, perimetro := rectangulo(5, 3)
	fmt.Printf("\nRectángulo 5x3: Área=%.1f, Perímetro=%.1f\n", area, perimetro)

	// Función variádica
	suma := sumarTodos(1, 2, 3, 4, 5)
	fmt.Printf("Suma de 1,2,3,4,5 = %d\n", suma)

	// Función como parámetro
	fmt.Println("\nFunciones como valores:")
	resultado = aplicarOperacion(10, 5, sumar)
	fmt.Printf("aplicarOperacion(10, 5, sumar) = %d\n", resultado)

	resultado = aplicarOperacion(10, 5, multiplicar)
	fmt.Printf("aplicarOperacion(10, 5, multiplicar) = %d\n", resultado)

	// Función anónima
	fmt.Println()
	ejemploFuncionAnonima()

	// Closure
	fmt.Println("\nClosure (contador):")
	incrementar := contador()
	fmt.Println(incrementar()) // 1
	fmt.Println(incrementar()) // 2
	fmt.Println(incrementar()) // 3

	// Otro contador independiente
	otroContador := contador()
	fmt.Println(otroContador()) // 1 (nuevo contador)
}

/*
CONCEPTOS CLAVE:

1. SINTAXIS: func nombre(params) retorno { cuerpo }

2. MÚLTIPLES RETORNOS: Muy común para retornar (valor, error)

3. FUNCIONES VARIÁDICAS: ...tipo permite número variable de args

4. FUNCIONES COMO VALORES: Puedes pasar funciones como parámetros

5. CLOSURES: Funciones que capturan variables de su entorno

6. NO HAY SOBRECARGA: No puedes tener dos funciones con el mismo nombre

7. CONVENCIÓN:
   - Funciones públicas empiezan con MAYÚSCULA (exportadas)
   - Funciones privadas empiezan con minúscula (dentro del paquete)

EJECUTAR:
go run 03_funciones.go
*/
