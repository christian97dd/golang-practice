package main

import (
	"fmt"
	"math"
)

// ========== INTERFACES ==========
// Una interfaz define un CONTRATO: un conjunto de métodos
// Cualquier tipo que implemente TODOS esos métodos satisface la interfaz
// En Go, las interfaces se implementan IMPLÍCITAMENTE (no necesitas declararlo)

// Interfaz simple con un método
type Figura interface {
	Area() float64
}

// Interfaz con múltiples métodos
type FiguraCompleta interface {
	Area() float64
	Perimetro() float64
}

// ========== STRUCTS QUE IMPLEMENTAN INTERFACES ==========

type Rectangulo struct {
	Ancho float64
	Alto  float64
}

// Rectangulo implementa la interfaz Figura (tiene método Area)
func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

// Rectangulo también implementa FiguraCompleta (tiene Area y Perimetro)
func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

type Circulo struct {
	Radio float64
}

// Circulo implementa Figura
func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

// Circulo implementa FiguraCompleta
func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

type Triangulo struct {
	Base   float64
	Altura float64
	Lado1  float64
	Lado2  float64
	Lado3  float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangulo) Perimetro() float64 {
	return t.Lado1 + t.Lado2 + t.Lado3
}

// ========== FUNCIONES QUE USAN INTERFACES ==========

// Esta función acepta CUALQUIER tipo que implemente Figura
func imprimirArea(f Figura) {
	fmt.Printf("Área: %.2f\n", f.Area())
}

// Acepta cualquier tipo que implemente FiguraCompleta
func imprimirInfo(f FiguraCompleta) {
	fmt.Printf("Área: %.2f, Perímetro: %.2f\n", f.Area(), f.Perimetro())
}

// Calcular área total de múltiples figuras
func areaTotal(figuras []Figura) float64 {
	total := 0.0
	for _, figura := range figuras {
		total += figura.Area()
	}
	return total
}

// ========== INTERFAZ VACÍA ==========
// interface{} acepta CUALQUIER tipo (como any en TypeScript)
// Desde Go 1.18 también puedes usar "any" (alias de interface{})

func imprimir(valor interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", valor, valor)
}

// ========== TYPE ASSERTION ==========
// Convertir una interfaz a un tipo concreto

func describir(i interface{}) {
	// Type assertion: i.(Tipo)
	// Retorna (valor, ok)
	if str, ok := i.(string); ok {
		fmt.Printf("Es un string de longitud %d: %s\n", len(str), str)
	} else if num, ok := i.(int); ok {
		fmt.Printf("Es un entero: %d\n", num)
	} else if f, ok := i.(Figura); ok {
		fmt.Printf("Es una Figura con área: %.2f\n", f.Area())
	} else {
		fmt.Printf("Tipo desconocido: %T\n", i)
	}
}

// ========== TYPE SWITCH ==========
// Forma más limpia de manejar múltiples tipos

func procesarValor(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("String: %s (longitud: %d)\n", v, len(v))
	case int:
		fmt.Printf("Int: %d (doble: %d)\n", v, v*2)
	case float64:
		fmt.Printf("Float64: %.2f\n", v)
	case Figura:
		fmt.Printf("Figura con área: %.2f\n", v.Area())
	case nil:
		fmt.Println("Es nil")
	default:
		fmt.Printf("Tipo no manejado: %T\n", v)
	}
}

// ========== INTERFACES PARA POLIMORFISMO ==========

type Animal interface {
	Sonido() string
	Mover() string
}

type Perro struct {
	Nombre string
}

func (p Perro) Sonido() string {
	return "Guau guau"
}

func (p Perro) Mover() string {
	return "Corriendo"
}

type Gato struct {
	Nombre string
}

func (g Gato) Sonido() string {
	return "Miau"
}

func (g Gato) Mover() string {
	return "Caminando sigilosamente"
}

type Pajaro struct {
	Nombre string
}

func (p Pajaro) Sonido() string {
	return "Pio pio"
}

func (p Pajaro) Mover() string {
	return "Volando"
}

func hacerSonido(a Animal) {
	fmt.Printf("%s hace: %s\n", obtenerNombre(a), a.Sonido())
}

func hacerMover(a Animal) {
	fmt.Printf("%s se mueve: %s\n", obtenerNombre(a), a.Mover())
}

// Función auxiliar que usa type assertion para obtener el nombre
func obtenerNombre(a Animal) string {
	switch v := a.(type) {
	case Perro:
		return v.Nombre
	case Gato:
		return v.Nombre
	case Pajaro:
		return v.Nombre
	default:
		return "Animal"
	}
}

// ========== INTERFACES ESTÁNDAR ==========
// Go tiene varias interfaces integradas muy útiles

// fmt.Stringer: permite personalizar cómo se imprime un tipo
type Producto struct {
	Nombre string
	Precio float64
}

// Implementar String() hace que el tipo implemente fmt.Stringer
func (p Producto) String() string {
	return fmt.Sprintf("%s - %.2f€", p.Nombre, p.Precio)
}

// ========== MAIN ==========

func main() {
	fmt.Println("=== INTERFACES EN GO ===\n")

	// ========== USO BÁSICO ==========
	fmt.Println("--- USO BÁSICO DE INTERFACES ---")

	rect := Rectangulo{Ancho: 10, Alto: 5}
	circ := Circulo{Radio: 7}
	tri := Triangulo{Base: 6, Altura: 4, Lado1: 5, Lado2: 5, Lado3: 6}

	// Pasar diferentes tipos a la misma función
	imprimirArea(rect)
	imprimirArea(circ)
	imprimirArea(tri)

	fmt.Println("\n--- Información completa ---")
	imprimirInfo(rect)
	imprimirInfo(circ)
	imprimirInfo(tri)

	// Slice de interfaces
	figuras := []Figura{rect, circ, tri}
	fmt.Printf("\nÁrea total de todas las figuras: %.2f\n", areaTotal(figuras))

	// ========== INTERFAZ VACÍA ==========
	fmt.Println("\n--- INTERFAZ VACÍA (any) ---")

	imprimir(42)
	imprimir("Hola Go")
	imprimir(3.14)
	imprimir(true)
	imprimir(rect)

	// ========== TYPE ASSERTION ==========
	fmt.Println("\n--- TYPE ASSERTION ---")

	describir("Hola mundo")
	describir(100)
	describir(rect)
	describir(true)

	// ========== TYPE SWITCH ==========
	fmt.Println("\n--- TYPE SWITCH ---")

	procesarValor("Go")
	procesarValor(42)
	procesarValor(3.14159)
	procesarValor(circ)
	procesarValor(nil)

	// ========== POLIMORFISMO ==========
	fmt.Println("\n--- POLIMORFISMO ---")

	perro := Perro{Nombre: "Max"}
	gato := Gato{Nombre: "Michi"}
	pajaro := Pajaro{Nombre: "Piolín"}

	animales := []Animal{perro, gato, pajaro}

	fmt.Println("Sonidos:")
	for _, animal := range animales {
		hacerSonido(animal)
	}

	fmt.Println("\nMovimientos:")
	for _, animal := range animales {
		hacerMover(animal)
	}

	// ========== fmt.Stringer ==========
	fmt.Println("\n--- fmt.Stringer ---")

	prod1 := Producto{Nombre: "Laptop", Precio: 999.99}
	prod2 := Producto{Nombre: "Mouse", Precio: 25.50}

	// fmt.Println usa automáticamente el método String()
	fmt.Println(prod1) // Usa String() personalizado
	fmt.Println(prod2)

	// Sin String(), imprimiría: {Laptop 999.99}

	// ========== COMPOSICIÓN DE INTERFACES ==========
	fmt.Println("\n--- COMPOSICIÓN DE INTERFACES ---")

	// Puedes componer interfaces
	type Nombrable interface {
		Nombre() string
	}

	type Guardable interface {
		Guardar() error
	}

	// Interfaz compuesta
	type Entidad interface {
		Nombrable
		Guardable
	}

	// Cualquier tipo que implemente Nombre() y Guardar() implementa Entidad
}

/*
CONCEPTOS CLAVE SOBRE INTERFACES:

1. DEFINICIÓN:
   type NombreInterfaz interface {
       Metodo1() tipoRetorno
       Metodo2(params) tipoRetorno
   }

2. IMPLEMENTACIÓN IMPLÍCITA:
   - No necesitas declarar "implements"
   - Si un tipo tiene TODOS los métodos, implementa la interfaz
   - Descubrimiento automático en tiempo de compilación

3. POLIMORFISMO:
   - Diferentes tipos pueden implementar la misma interfaz
   - Funciones aceptan la interfaz, no tipos concretos
   - Permite código flexible y reutilizable

4. INTERFAZ VACÍA (interface{} o any):
   - Acepta CUALQUIER tipo
   - Útil para funciones genéricas
   - Requiere type assertion para usar

5. TYPE ASSERTION:
   - valor, ok := interfaz.(Tipo)
   - Convierte interfaz a tipo concreto
   - Retorna valor y booleano (éxito)

6. TYPE SWITCH:
   - switch v := i.(type) { case Tipo: ... }
   - Forma limpia de manejar múltiples tipos

7. INTERFACES ESTÁNDAR:
   - fmt.Stringer: String() string
   - io.Reader: Read([]byte) (int, error)
   - io.Writer: Write([]byte) (int, error)
   - error: Error() string

8. BUENAS PRÁCTICAS:
   - Interfaces pequeñas (1-3 métodos)
   - Definir en el paquete consumidor, no productor
   - Nombres descriptivos (ej: Reader, Writer, Stringer)
   - Sufijo -er común (Sorter, Runner, etc.)

9. VENTAJAS:
   - Acoplamiento débil
   - Testeo fácil (mocks)
   - Código flexible
   - Composición sobre herencia

DIFERENCIAS CON OTROS LENGUAJES:
- Java/C#: Implementación explícita (implements)
- Go: Implementación implícita (duck typing estático)
- JavaScript: Duck typing dinámico
- TypeScript: Interfaces solo en tiempo de compilación

EJECUTAR:
go run 09_interfaces.go
*/
