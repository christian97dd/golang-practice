package main

import "fmt"

// ========== PUNTEROS EN GO ==========
// Un puntero guarda la DIRECCIÓN DE MEMORIA de una variable
// Permite modificar valores directamente en memoria (pasar por referencia)

func main() {
	fmt.Println("=== PUNTEROS ===\n")

	// ========== DECLARACIÓN Y USO BÁSICO ==========
	fmt.Println("--- DECLARACIÓN BÁSICA ---")

	// Variable normal
	x := 42
	fmt.Printf("x = %d\n", x)
	fmt.Printf("Dirección de x: %p\n", &x) // & obtiene la dirección

	// Puntero a x
	var p *int    // p es un puntero a int (puede guardar la dirección de un int)
	p = &x        // & obtiene la dirección de x
	fmt.Printf("p (dirección que guarda) = %p\n", p)
	fmt.Printf("*p (valor al que apunta) = %d\n", *p) // * desreferencia (obtiene el valor)

	// Modificar a través del puntero
	*p = 100 // Cambia el valor de x a través del puntero
	fmt.Printf("Después de *p = 100:\n")
	fmt.Printf("x = %d\n", x)   // x ahora es 100
	fmt.Printf("*p = %d\n", *p) // *p también es 100

	// ========== OPERADORES ==========
	fmt.Println("\n--- OPERADORES ---")

	numero := 10
	puntero := &numero

	fmt.Printf("numero = %d\n", numero)
	fmt.Printf("&numero (dirección) = %p\n", &numero)
	fmt.Printf("puntero = %p\n", puntero)
	fmt.Printf("*puntero (desreferencia) = %d\n", *puntero)

	// & : dirección de (address of)
	// * : desreferenciar (dereferencing) cuando se usa con un puntero
	// * : declarar tipo puntero cuando se usa en declaración

	// ========== PUNTERO NIL ==========
	fmt.Println("\n--- PUNTERO NIL ---")

	var ptr *int // Puntero sin inicializar es nil
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Printf("¿Es nil? %t\n", ptr == nil)

	// PELIGRO: Desreferenciar un puntero nil causa panic
	// fmt.Println(*ptr) // ¡ESTO CAUSARÍA PANIC!

	// Siempre verificar antes de desreferenciar
	if ptr != nil {
		fmt.Println("Valor:", *ptr)
	} else {
		fmt.Println("El puntero es nil, no se puede desreferenciar")
	}

	// ========== PASAR POR VALOR vs REFERENCIA ==========
	fmt.Println("\n--- PASAR POR VALOR vs REFERENCIA ---")

	// Por valor (se copia)
	a := 5
	fmt.Printf("Antes de duplicarValor: a = %d\n", a)
	duplicarValor(a)
	fmt.Printf("Después de duplicarValor: a = %d (sin cambios)\n", a)

	// Por referencia (con puntero)
	b := 5
	fmt.Printf("\nAntes de duplicarReferencia: b = %d\n", b)
	duplicarReferencia(&b) // Pasamos la dirección de b
	fmt.Printf("Después de duplicarReferencia: b = %d (cambiado)\n", b)

	// ========== PUNTEROS CON STRUCTS ==========
	fmt.Println("\n--- PUNTEROS CON STRUCTS ---")

	type Persona struct {
		Nombre string
		Edad   int
	}

	// Crear struct normal
	persona1 := Persona{Nombre: "Ana", Edad: 25}
	fmt.Printf("persona1: %+v\n", persona1)

	// Crear puntero a struct
	persona2 := &Persona{Nombre: "Carlos", Edad: 30}
	fmt.Printf("persona2: %+v\n", persona2) // Go imprime el struct, no la dirección

	// Acceder a campos con puntero
	// Go permite usar . directamente (azúcar sintáctica)
	fmt.Printf("Nombre: %s\n", persona2.Nombre) // Equivalente a (*persona2).Nombre
	persona2.Edad = 31
	fmt.Printf("persona2 después: %+v\n", persona2)

	// Modificar struct con función
	modificarPersonaValor(persona1)
	fmt.Printf("persona1 después de modificarPersonaValor: %+v (sin cambios)\n", persona1)

	modificarPersonaPuntero(persona2)
	fmt.Printf("persona2 después de modificarPersonaPuntero: %+v (cambiado)\n", persona2)

	// ========== NEW ==========
	fmt.Println("\n--- NEW ---")

	// new() crea un puntero a un valor cero
	punteroInt := new(int)
	fmt.Printf("*punteroInt (valor inicial): %d\n", *punteroInt) // 0
	*punteroInt = 50
	fmt.Printf("*punteroInt (después): %d\n", *punteroInt)

	punteroPersona := new(Persona)
	fmt.Printf("punteroPersona: %+v\n", punteroPersona) // {Nombre: Edad:0}
	punteroPersona.Nombre = "Luis"
	punteroPersona.Edad = 28
	fmt.Printf("punteroPersona: %+v\n", punteroPersona)

	// ========== PUNTEROS CON SLICES Y MAPS ==========
	fmt.Println("\n--- PUNTEROS CON SLICES Y MAPS ---")

	// IMPORTANTE: Slices y maps YA son referencias
	// No necesitas usar punteros con ellos (normalmente)

	slice := []int{1, 2, 3}
	fmt.Printf("Slice original: %v\n", slice)
	modificarSlice(slice) // No necesitas &slice
	fmt.Printf("Slice después: %v (modificado)\n", slice)

	mapa := make(map[string]int)
	mapa["a"] = 1
	fmt.Printf("\nMapa original: %v\n", mapa)
	modificarMapa(mapa) // No necesitas &mapa
	fmt.Printf("Mapa después: %v (modificado)\n", mapa)

	// ========== COMPARACIÓN DE PUNTEROS ==========
	fmt.Println("\n--- COMPARACIÓN ---")

	n1 := 10
	n2 := 10
	p1 := &n1
	p2 := &n2
	p3 := &n1

	fmt.Printf("p1 == p2: %t (diferentes direcciones)\n", p1 == p2)
	fmt.Printf("p1 == p3: %t (misma dirección)\n", p1 == p3)
	fmt.Printf("*p1 == *p2: %t (mismos valores)\n", *p1 == *p2)

	// ========== CUÁNDO USAR PUNTEROS ==========
	fmt.Println("\n--- CUÁNDO USAR PUNTEROS ---")

	type GranStruct struct {
		Datos [1000]int // Struct grande
	}

	// Para structs grandes, usa punteros para evitar copias costosas
	granStruct := &GranStruct{}
	procesarGranStruct(granStruct) // Eficiente

	// Para structs pequeños, el valor está bien
	type PequenoStruct struct {
		X int
		Y int
	}

	pequenoStruct := PequenoStruct{X: 1, Y: 2}
	procesarPequenoStruct(pequenoStruct) // Ok pasar por valor
}

// ========== FUNCIONES DE EJEMPLO ==========

// Recibe una copia del valor
func duplicarValor(n int) {
	n = n * 2
	fmt.Printf("  Dentro de duplicarValor: n = %d\n", n)
	// Los cambios no afectan la variable original
}

// Recibe un puntero (referencia)
func duplicarReferencia(n *int) {
	*n = *n * 2
	fmt.Printf("  Dentro de duplicarReferencia: *n = %d\n", *n)
	// Los cambios SÍ afectan la variable original
}

// Recibe struct por valor (se copia)
func modificarPersonaValor(p Persona) {
	p.Edad = 100
	fmt.Printf("  Dentro de modificarPersonaValor: %+v\n", p)
}

// Recibe puntero a struct (referencia)
func modificarPersonaPuntero(p *Persona) {
	p.Edad = 100 // Go automáticamente desreferencia
	fmt.Printf("  Dentro de modificarPersonaPuntero: %+v\n", p)
}

type Persona struct {
	Nombre string
	Edad   int
}

// Slices se pasan por referencia (son referencias internamente)
func modificarSlice(s []int) {
	s[0] = 999 // Modifica el slice original
}

// Maps se pasan por referencia (son referencias internamente)
func modificarMapa(m map[string]int) {
	m["b"] = 2 // Modifica el mapa original
}

type GranStruct struct {
	Datos [1000]int
}

type PequenoStruct struct {
	X int
	Y int
}

func procesarGranStruct(g *GranStruct) {
	// Procesar sin copiar 1000 ints
	g.Datos[0] = 1
}

func procesarPequenoStruct(p PequenoStruct) {
	// Copiar 2 ints es barato
	_ = p.X + p.Y
}

/*
CONCEPTOS CLAVE SOBRE PUNTEROS:

1. DEFINICIÓN:
   - Puntero: variable que guarda una DIRECCIÓN DE MEMORIA
   - *T: tipo puntero a T
   - &x: dirección de x (address of)
   - *p: valor al que apunta p (dereference)

2. SINTAXIS:
   var p *int      // Declarar puntero a int
   p = &x          // Asignar dirección de x
   *p = 10         // Modificar valor al que apunta

3. NIL:
   - Puntero sin inicializar es nil
   - Desreferenciar nil causa PANIC
   - Siempre verificar: if p != nil { ... }

4. PASAR POR VALOR vs REFERENCIA:
   VALOR:
   - func(x int): se copia
   - Cambios no afectan original
   - Usa para tipos pequeños

   REFERENCIA:
   - func(x *int): pasa dirección
   - Cambios SÍ afectan original
   - Usa para modificar o tipos grandes

5. STRUCTS:
   - &Struct{}: crea puntero a struct
   - p.Campo: Go desreferencia automáticamente
   - No necesitas (*p).Campo

6. SLICES Y MAPS:
   - YA son referencias (no necesitas punteros)
   - Pasar slice/map modifica el original

7. NEW:
   - new(T): crea puntero a valor cero de T
   - Alternativa: &T{}

8. CUÁNDO USAR PUNTEROS:
   ✓ Modificar parámetros
   ✓ Structs grandes (evitar copias)
   ✓ Compartir estado
   ✓ Métodos que modifican (receptor de puntero)

   ✗ Tipos básicos pequeños (int, bool)
   ✗ Slices y maps (ya son referencias)
   ✗ Cuando no necesitas modificar

9. DIFERENCIAS CON C/C++:
   - NO hay aritmética de punteros
   - NO hay punteros void*
   - Más seguros (garbage collection)
   - Sintaxis más simple

10. SEGURIDAD:
    - Go maneja la memoria automáticamente
    - No necesitas free/delete
    - Garbage collector limpia memoria no usada

EJECUTAR:
go run 10_punteros.go
*/
