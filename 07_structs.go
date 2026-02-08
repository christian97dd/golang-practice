package main

import "fmt"

// ========== DEFINIR STRUCTS ==========
// Los structs son como clases en otros lenguajes
// Agrupan datos relacionados (similar a objetos en JavaScript)

// Struct simple
type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

// Struct con campos de diferentes tipos
type Direccion struct {
	Calle  string
	Numero int
	Ciudad string
	CP     string
}

// Struct con otro struct (composición)
type Empleado struct {
	Nombre    string
	Edad      int
	Direccion Direccion // Struct anidado
	Salario   float64
}

// Struct con campos opcionales usando punteros
type Usuario struct {
	ID       int
	Username string
	Email    *string // Puntero: puede ser nil (opcional)
	Activo   bool
}

// ========== STRUCT ANÓNIMO ==========
// Puedes crear structs sin nombre para uso temporal

func main() {
	fmt.Println("=== STRUCTS ===\n")

	// ========== CREAR INSTANCIAS ==========
	fmt.Println("--- CREAR STRUCTS ---")

	// Forma 1: Valores cero
	var p1 Persona
	fmt.Printf("Persona vacía: %+v\n", p1) // %+v muestra nombres de campos

	// Forma 2: Literal con nombres de campos
	p2 := Persona{
		Nombre: "Ana",
		Edad:   25,
		Email:  "ana@email.com",
	}
	fmt.Printf("Persona 2: %+v\n", p2)

	// Forma 3: Literal sin nombres (mismo orden que definición)
	p3 := Persona{"Carlos", 30, "carlos@email.com"}
	fmt.Printf("Persona 3: %+v\n", p3)

	// Forma 4: Valores parciales (resto son valores cero)
	p4 := Persona{
		Nombre: "Luis",
		// Edad y Email tendrán valores cero
	}
	fmt.Printf("Persona 4: %+v\n", p4)

	// ========== ACCEDER A CAMPOS ==========
	fmt.Println("\n--- ACCEDER Y MODIFICAR CAMPOS ---")

	fmt.Println("Nombre:", p2.Nombre)
	fmt.Println("Edad:", p2.Edad)

	// Modificar campos
	p2.Edad = 26
	p2.Email = "ana.nueva@email.com"
	fmt.Printf("Persona modificada: %+v\n", p2)

	// ========== STRUCTS ANIDADOS ==========
	fmt.Println("\n--- STRUCTS ANIDADOS ---")

	empleado := Empleado{
		Nombre: "María",
		Edad:   28,
		Direccion: Direccion{
			Calle:  "Gran Vía",
			Numero: 45,
			Ciudad: "Madrid",
			CP:     "28013",
		},
		Salario: 35000.50,
	}

	fmt.Printf("Empleado: %+v\n", empleado)
	fmt.Printf("Ciudad: %s\n", empleado.Direccion.Ciudad)
	fmt.Printf("Calle: %s, %d\n", empleado.Direccion.Calle, empleado.Direccion.Numero)

	// ========== PUNTEROS A STRUCTS ==========
	fmt.Println("\n--- PUNTEROS A STRUCTS ---")

	// Crear puntero a struct
	punteroPersona := &Persona{
		Nombre: "Pedro",
		Edad:   35,
		Email:  "pedro@email.com",
	}

	// Go automáticamente desreferencia
	fmt.Println("Nombre:", punteroPersona.Nombre) // No necesitas (*punteroPersona).Nombre
	punteroPersona.Edad = 36
	fmt.Printf("Persona via puntero: %+v\n", punteroPersona)

	// ========== STRUCTS ANÓNIMOS ==========
	fmt.Println("\n--- STRUCTS ANÓNIMOS ---")

	// Útil para datos temporales
	punto := struct {
		X int
		Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Printf("Punto: %+v\n", punto)

	// ========== COMPARACIÓN ==========
	fmt.Println("\n--- COMPARACIÓN ---")

	// Los structs se pueden comparar si todos sus campos son comparables
	per1 := Persona{"Ana", 25, "ana@email.com"}
	per2 := Persona{"Ana", 25, "ana@email.com"}
	per3 := Persona{"Carlos", 25, "carlos@email.com"}

	fmt.Println("per1 == per2:", per1 == per2) // true
	fmt.Println("per1 == per3:", per1 == per3) // false

	// ========== COPIAR STRUCTS ==========
	fmt.Println("\n--- COPIAR STRUCTS ---")

	original := Persona{"Original", 30, "original@email.com"}
	copia := original // Se copia por valor

	copia.Nombre = "Copia Modificada"

	fmt.Printf("Original: %+v\n", original) // No cambia
	fmt.Printf("Copia: %+v\n", copia)       // Cambiada

	// ========== TAGS (etiquetas) ==========
	fmt.Println("\n--- TAGS ---")

	// Los tags son metadatos útiles para JSON, validación, etc.
	type Producto struct {
		ID     int     `json:"id"`
		Nombre string  `json:"nombre"`
		Precio float64 `json:"precio"`
	}

	prod := Producto{ID: 1, Nombre: "Laptop", Precio: 999.99}
	fmt.Printf("Producto: %+v\n", prod)
	// Las tags se usan principalmente con encoding/json

	// ========== FUNCIONES CON STRUCTS ==========
	fmt.Println("\n--- FUNCIONES CON STRUCTS ---")

	persona := Persona{"Roberto", 40, "roberto@email.com"}

	// Pasar por valor (se copia)
	imprimirPersona(persona)

	// Pasar por referencia (con puntero)
	cumpleanios(&persona) // Modifica el original
	fmt.Printf("Después del cumpleaños: %+v\n", persona)

	// Función que retorna struct
	nueva := crearPersona("Laura", 22, "laura@email.com")
	fmt.Printf("Nueva persona: %+v\n", nueva)

	// ========== SLICE DE STRUCTS ==========
	fmt.Println("\n--- SLICE DE STRUCTS ---")

	personas := []Persona{
		{"Ana", 25, "ana@email.com"},
		{"Carlos", 30, "carlos@email.com"},
		{"Luis", 28, "luis@email.com"},
	}

	fmt.Println("Lista de personas:")
	for i, p := range personas {
		fmt.Printf("  %d. %s (%d años)\n", i+1, p.Nombre, p.Edad)
	}

	// ========== MAP DE STRUCTS ==========
	fmt.Println("\n--- MAP DE STRUCTS ---")

	usuarios := map[int]Persona{
		1: {"Ana", 25, "ana@email.com"},
		2: {"Carlos", 30, "carlos@email.com"},
		3: {"Luis", 28, "luis@email.com"},
	}

	fmt.Println("Usuarios:")
	for id, usuario := range usuarios {
		fmt.Printf("  ID %d: %s\n", id, usuario.Nombre)
	}
}

// Función que recibe struct por valor (se copia)
func imprimirPersona(p Persona) {
	fmt.Printf("Imprimir: %s tiene %d años\n", p.Nombre, p.Edad)
}

// Función que recibe struct por referencia (puntero)
// Puede modificar el struct original
func cumpleanios(p *Persona) {
	p.Edad++
	fmt.Printf("¡Feliz cumpleaños %s! Ahora tienes %d años\n", p.Nombre, p.Edad)
}

// Función que retorna un struct
func crearPersona(nombre string, edad int, email string) Persona {
	return Persona{
		Nombre: nombre,
		Edad:   edad,
		Email:  email,
	}
}

/*
CONCEPTOS CLAVE SOBRE STRUCTS:

1. DEFINICIÓN:
   type NombreStruct struct {
       Campo1 Tipo1
       Campo2 Tipo2
   }

2. CREAR INSTANCIAS:
   - var s Struct                    -> valores cero
   - s := Struct{campo: valor}       -> con nombres
   - s := Struct{val1, val2}         -> sin nombres (orden)
   - s := &Struct{...}               -> puntero a struct

3. ACCESO:
   - struct.campo                    -> acceder/modificar
   - puntero.campo                   -> Go desreferencia automáticamente

4. CARACTERÍSTICAS:
   - Tipos por VALOR (se copian)
   - Comparables si todos los campos son comparables
   - No hay herencia (usar composición)
   - Campos públicos: MAYÚSCULA
   - Campos privados: minúscula

5. COMPOSICIÓN:
   - Incluir structs dentro de otros structs
   - Embedding (incrustación) para reutilizar código

6. PUNTEROS:
   - Pasar &struct para modificar el original
   - Más eficiente para structs grandes

7. TAGS:
   - Metadatos para serialización (JSON, DB, etc.)
   - `json:"nombre"` `db:"nombre"`

DIFERENCIAS CON JS/TS:
- No hay clases, solo structs (datos)
- Métodos se definen fuera del struct (ver 08_metodos.go)
- No hay constructores (usar funciones factorías)
- No hay this/self automático

EJECUTAR:
go run 07_structs.go
*/
