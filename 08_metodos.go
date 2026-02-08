package main

import (
	"fmt"
	"math"
)

// ========== STRUCTS ==========

type Rectangulo struct {
	Ancho  float64
	Alto   float64
}

type Circulo struct {
	Radio float64
}

type Persona struct {
	Nombre string
	Edad   int
}

type Contador struct {
	valor int
}

// ========== MÉTODOS ==========
// Los métodos son funciones asociadas a un tipo (struct)
// Se definen FUERA del struct, con un "receptor"

// Método con receptor de valor (recibe una COPIA)
// Sintaxis: func (receptor TipoReceptor) nombreMetodo() tipoRetorno
func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

// Otro método para Rectangulo
func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

// Método que retorna múltiples valores
func (r Rectangulo) Dimensiones() (float64, float64) {
	return r.Ancho, r.Alto
}

// Método para Circulo
func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

// ========== RECEPTOR DE VALOR vs PUNTERO ==========

// Receptor de VALOR: recibe una COPIA
// NO puede modificar el struct original
func (p Persona) Saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años\n", p.Nombre, p.Edad)
}

// Intento de modificar con receptor de valor (NO funciona)
func (p Persona) CumpleaniosIncorrecto() {
	p.Edad++ // Modifica la COPIA, no el original
	fmt.Printf("Método con valor: ahora tengo %d años (no cambia el original)\n", p.Edad)
}

// Receptor de PUNTERO: recibe una REFERENCIA
// SÍ puede modificar el struct original
func (p *Persona) Cumpleanios() {
	p.Edad++ // Modifica el original
	fmt.Printf("¡Feliz cumpleaños! Ahora tengo %d años\n", p.Edad)
}

func (p *Persona) CambiarNombre(nuevoNombre string) {
	p.Nombre = nuevoNombre
}

// ========== MÉTODOS QUE MODIFICAN ESTADO ==========

func (c *Contador) Incrementar() {
	c.valor++
}

func (c *Contador) Decrementar() {
	c.valor--
}

func (c *Contador) Valor() int {
	return c.valor
}

func (c *Contador) Reset() {
	c.valor = 0
}

// ========== MÉTODO CON PARÁMETROS ==========

func (r Rectangulo) Escalar(factor float64) Rectangulo {
	return Rectangulo{
		Ancho: r.Ancho * factor,
		Alto:  r.Alto * factor,
	}
}

func (c *Circulo) Agrandar(incremento float64) {
	c.Radio += incremento
}

// ========== MÉTODOS EN TIPOS BÁSICOS ==========
// Puedes definir métodos en tipos personalizados basados en tipos básicos

type Temperatura float64

func (t Temperatura) Celsius() float64 {
	return float64(t)
}

func (t Temperatura) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (t Temperatura) Kelvin() float64 {
	return float64(t) + 273.15
}

// ========== FUNCIÓN MAIN ==========

func main() {
	fmt.Println("=== MÉTODOS EN GO ===\n")

	// ========== MÉTODOS BÁSICOS ==========
	fmt.Println("--- MÉTODOS BÁSICOS ---")

	rect := Rectangulo{Ancho: 10, Alto: 5}
	fmt.Printf("Rectángulo: %+v\n", rect)
	fmt.Printf("Área: %.2f\n", rect.Area())
	fmt.Printf("Perímetro: %.2f\n", rect.Perimetro())

	ancho, alto := rect.Dimensiones()
	fmt.Printf("Dimensiones: %.2f x %.2f\n", ancho, alto)

	circ := Circulo{Radio: 7}
	fmt.Printf("\nCírculo: %+v\n", circ)
	fmt.Printf("Área: %.2f\n", circ.Area())
	fmt.Printf("Perímetro: %.2f\n", circ.Perimetro())

	// ========== RECEPTOR VALOR vs PUNTERO ==========
	fmt.Println("\n--- RECEPTOR: VALOR vs PUNTERO ---")

	persona := Persona{Nombre: "Ana", Edad: 25}
	fmt.Printf("Persona inicial: %+v\n", persona)

	persona.Saludar()

	// Método con receptor de valor (no modifica)
	persona.CumpleaniosIncorrecto()
	fmt.Printf("Después de CumpleaniosIncorrecto: %+v\n", persona) // Sigue 25

	// Método con receptor de puntero (sí modifica)
	persona.Cumpleanios()
	fmt.Printf("Después de Cumpleanios: %+v\n", persona) // Ahora 26

	// Go automáticamente toma la dirección si es necesario
	persona.CambiarNombre("Ana María") // Go convierte a (&persona).CambiarNombre()
	fmt.Printf("Después de cambiar nombre: %+v\n", persona)

	// También funciona al revés: con puntero llamar método de valor
	punteroPersona := &Persona{Nombre: "Carlos", Edad: 30}
	punteroPersona.Saludar() // Go convierte a (*punteroPersona).Saludar()

	// ========== CONTADOR CON ESTADO ==========
	fmt.Println("\n--- CONTADOR ---")

	cont := Contador{valor: 0}
	fmt.Printf("Valor inicial: %d\n", cont.Valor())

	cont.Incrementar()
	cont.Incrementar()
	cont.Incrementar()
	fmt.Printf("Después de 3 incrementos: %d\n", cont.Valor())

	cont.Decrementar()
	fmt.Printf("Después de 1 decremento: %d\n", cont.Valor())

	cont.Reset()
	fmt.Printf("Después de reset: %d\n", cont.Valor())

	// ========== MÉTODOS CON PARÁMETROS ==========
	fmt.Println("\n--- MÉTODOS CON PARÁMETROS ---")

	rect2 := Rectangulo{Ancho: 5, Alto: 3}
	fmt.Printf("Rectángulo original: %+v, Área: %.2f\n", rect2, rect2.Area())

	// Escalar retorna un NUEVO rectángulo (receptor de valor)
	rect3 := rect2.Escalar(2)
	fmt.Printf("Rectángulo escalado x2: %+v, Área: %.2f\n", rect3, rect3.Area())
	fmt.Printf("Original sin cambios: %+v\n", rect2)

	circ2 := Circulo{Radio: 5}
	fmt.Printf("\nCírculo original: %+v, Área: %.2f\n", circ2, circ2.Area())

	// Agrandar modifica el original (receptor de puntero)
	circ2.Agrandar(3)
	fmt.Printf("Círculo agrandado: %+v, Área: %.2f\n", circ2, circ2.Area())

	// ========== MÉTODOS EN TIPOS PERSONALIZADOS ==========
	fmt.Println("\n--- MÉTODOS EN TIPOS PERSONALIZADOS ---")

	temp := Temperatura(25.0)
	fmt.Printf("Temperatura: %.1f°C\n", temp.Celsius())
	fmt.Printf("En Fahrenheit: %.1f°F\n", temp.Fahrenheit())
	fmt.Printf("En Kelvin: %.1fK\n", temp.Kelvin())

	// ========== ENCADENAMIENTO DE MÉTODOS ==========
	fmt.Println("\n--- ENCADENAMIENTO ---")

	// Para encadenar, los métodos deben retornar el mismo tipo
	type Constructor struct {
		valor int
	}

	// Estos métodos retornan *Constructor para permitir encadenamiento
	func (c *Constructor) Sumar(n int) *Constructor {
		c.valor += n
		return c
	}

	func (c *Constructor) Multiplicar(n int) *Constructor {
		c.valor *= n
		return c
	}

	func (c *Constructor) Obtener() int {
		return c.valor
	}

	builder := &Constructor{valor: 10}
	resultado := builder.Sumar(5).Multiplicar(2).Obtener()
	fmt.Printf("(10 + 5) * 2 = %d\n", resultado)
}

/*
CONCEPTOS CLAVE SOBRE MÉTODOS:

1. SINTAXIS:
   func (receptor Tipo) NombreMetodo(params) retorno {
       // código
   }

2. RECEPTOR DE VALOR vs PUNTERO:

   VALOR (r Tipo):
   - Recibe una COPIA del struct
   - NO puede modificar el original
   - Más seguro (no hay efectos secundarios)
   - Usa cuando: solo lees datos, struct pequeño

   PUNTERO (r *Tipo):
   - Recibe una REFERENCIA al struct
   - SÍ puede modificar el original
   - Más eficiente (no copia)
   - Usa cuando: necesitas modificar, struct grande

3. CONVENCIÓN:
   - Si UN método usa puntero, TODOS deberían usarlo
   - Go convierte automáticamente: valor.MetodoPuntero() funciona
   - Tipos pequeños (int, bool): receptor de valor
   - Tipos grandes o mutables: receptor de puntero

4. DIFERENCIAS CON OOP:
   - No hay clases, solo tipos con métodos
   - No hay this/self (usas el nombre del receptor)
   - No hay herencia (usa composición)
   - Métodos definidos FUERA del struct

5. VENTAJAS:
   - Sintaxis clara: tipo.Metodo()
   - Separación entre datos (struct) y comportamiento (métodos)
   - Puedes agregar métodos a cualquier tipo (incluso básicos)

6. CUÁNDO USAR PUNTERO:
   - Modificar el struct
   - Struct grande (evitar copia)
   - Mantener consistencia (si otros métodos usan puntero)

EJECUTAR:
go run 08_metodos.go
*/
