package main

import "fmt"

func main() {
	fmt.Println("=== MAPS (Diccionarios/Objetos) ===\n")

	// ========== CREAR MAPS ==========
	// Maps son colecciones de pares clave-valor (como objetos en JS)
	// Sintaxis: map[tipoClave]tipoValor

	// Forma 1: Declarar map vacío (es nil)
	var edades map[string]int
	fmt.Println("Map vacío:", edades)
	fmt.Println("¿Es nil?:", edades == nil) // true
	// NO puedes agregar elementos a un map nil

	// Forma 2: Crear con make (la más común)
	edades = make(map[string]int)
	fmt.Println("Map con make:", edades) // map[]

	// Forma 3: Map literal (declarar e inicializar)
	capitales := map[string]string{
		"España":    "Madrid",
		"Francia":   "París",
		"Italia":    "Roma",
		"Alemania":  "Berlín",
	}
	fmt.Println("Capitales:", capitales)

	// ========== AGREGAR Y MODIFICAR ==========
	fmt.Println("\n--- AGREGAR Y MODIFICAR ---")

	// Agregar elementos
	edades["Ana"] = 25
	edades["Carlos"] = 30
	edades["Luis"] = 28
	fmt.Println("Edades:", edades)

	// Modificar (misma sintaxis)
	edades["Ana"] = 26 // Cambia el valor
	fmt.Println("Ana ahora tiene:", edades["Ana"])

	// ========== ACCEDER A VALORES ==========
	fmt.Println("\n--- ACCEDER A VALORES ---")

	// Acceder directamente
	fmt.Println("Edad de Carlos:", edades["Carlos"])

	// Si la clave no existe, retorna el valor cero del tipo
	fmt.Println("Edad de Pedro (no existe):", edades["Pedro"]) // 0

	// Verificar si una clave existe (muy importante)
	edad, existe := edades["Ana"]
	if existe {
		fmt.Printf("Ana existe y tiene %d años\n", edad)
	}

	// Verificar clave inexistente
	edad, existe = edades["María"]
	if !existe {
		fmt.Println("María no está en el map")
	}

	// Forma corta
	if edad, ok := edades["Carlos"]; ok {
		fmt.Printf("Carlos tiene %d años\n", edad)
	}

	// ========== ELIMINAR ELEMENTOS ==========
	fmt.Println("\n--- ELIMINAR ---")

	fmt.Println("Antes de eliminar:", edades)
	delete(edades, "Luis") // Elimina la clave "Luis"
	fmt.Println("Después de eliminar Luis:", edades)

	// Eliminar clave inexistente no causa error
	delete(edades, "Inexistente") // No pasa nada

	// ========== ITERAR SOBRE MAPS ==========
	fmt.Println("\n--- ITERAR ---")

	for pais, capital := range capitales {
		fmt.Printf("  La capital de %s es %s\n", pais, capital)
	}

	// Solo claves
	fmt.Print("Países: ")
	for pais := range capitales {
		fmt.Print(pais, " ")
	}
	fmt.Println()

	// Solo valores (usar _ para ignorar clave)
	fmt.Print("Capitales: ")
	for _, capital := range capitales {
		fmt.Print(capital, " ")
	}
	fmt.Println()

	// ========== LONGITUD ==========
	fmt.Println("\n--- LONGITUD ---")
	fmt.Printf("Número de países: %d\n", len(capitales))

	// ========== MAPS COMPLEJOS ==========
	fmt.Println("\n--- MAPS COMPLEJOS ---")

	// Map de slices
	familias := map[string][]string{
		"García":    {"Juan", "María", "Pedro"},
		"López":     {"Ana", "Luis"},
		"Martínez":  {"Carlos", "Laura", "José", "Elena"},
	}

	fmt.Println("Familias:")
	for apellido, miembros := range familias {
		fmt.Printf("  Familia %s: %v\n", apellido, miembros)
	}

	// Map de maps
	estudiantes := map[string]map[string]interface{}{
		"estudiante1": {
			"nombre": "Ana",
			"edad":   20,
			"nota":   8.5,
		},
		"estudiante2": {
			"nombre": "Carlos",
			"edad":   22,
			"nota":   9.0,
		},
	}

	fmt.Println("\nEstudiantes:")
	for id, datos := range estudiantes {
		fmt.Printf("  %s: %v\n", id, datos)
	}

	// ========== MAPS Y FUNCIONES ==========
	fmt.Println("\n--- FUNCIONES CON MAPS ---")

	productos := map[string]float64{
		"Pan":     1.50,
		"Leche":   2.30,
		"Huevos":  3.20,
		"Queso":   5.80,
	}

	fmt.Println("Productos:", productos)
	fmt.Printf("Producto más caro: %s a %.2f€\n", productoMasCaro(productos))

	// Modificar map en función
	aplicarDescuento(productos, 0.10) // 10% descuento
	fmt.Println("Con descuento:", productos)

	// ========== CASOS DE USO ==========
	fmt.Println("\n--- CASOS DE USO ---")

	// Contar ocurrencias
	palabras := []string{"hola", "mundo", "hola", "go", "mundo", "hola"}
	conteo := contarPalabras(palabras)
	fmt.Println("Conteo de palabras:", conteo)

	// Agrupar por categoría
	personas := []string{"Ana", "Alberto", "Carlos", "Camila", "Luis", "Laura"}
	porInicial := agruparPorInicial(personas)
	fmt.Println("\nAgrupado por inicial:")
	for inicial, nombres := range porInicial {
		fmt.Printf("  %c: %v\n", inicial, nombres)
	}
}

// Función que retorna la clave del valor máximo
func productoMasCaro(precios map[string]float64) (string, float64) {
	var maxProducto string
	var maxPrecio float64

	// Inicializar con el primer elemento
	primera := true
	for producto, precio := range precios {
		if primera || precio > maxPrecio {
			maxProducto = producto
			maxPrecio = precio
			primera = false
		}
	}

	return maxProducto, maxPrecio
}

// Maps se pasan por referencia
// Modificar el map dentro de la función afecta al original
func aplicarDescuento(precios map[string]float64, descuento float64) {
	for producto := range precios {
		precios[producto] *= (1 - descuento)
	}
}

// Contar ocurrencias de elementos
func contarPalabras(palabras []string) map[string]int {
	conteo := make(map[string]int)
	for _, palabra := range palabras {
		conteo[palabra]++ // Incrementa (valor cero es 0)
	}
	return conteo
}

// Agrupar elementos por criterio
func agruparPorInicial(nombres []string) map[rune][]string {
	grupos := make(map[rune][]string)
	for _, nombre := range nombres {
		inicial := rune(nombre[0]) // Primera letra
		grupos[inicial] = append(grupos[inicial], nombre)
	}
	return grupos
}

/*
CONCEPTOS CLAVE SOBRE MAPS:

1. DECLARACIÓN:
   - var m map[K]V           -> nil map (no usar directamente)
   - m := make(map[K]V)      -> map vacío (usar este)
   - m := map[K]V{...}       -> map con valores iniciales

2. OPERACIONES:
   - Agregar/Modificar: m[clave] = valor
   - Acceder: valor := m[clave]
   - Verificar existencia: valor, ok := m[clave]
   - Eliminar: delete(m, clave)
   - Longitud: len(m)

3. ITERACIÓN:
   - for k, v := range m { ... }  -> clave y valor
   - for k := range m { ... }     -> solo claves
   - for _, v := range m { ... }  -> solo valores

4. CARACTERÍSTICAS:
   - Desordenados (orden de iteración no garantizado)
   - Tipos de clave: comparables (int, string, struct sin slices)
   - Tipos de valor: cualquiera
   - Se pasan por REFERENCIA a funciones
   - NO son seguros para concurrencia (usar sync.Map)

5. VALOR CERO:
   - Si la clave no existe, retorna el valor cero del tipo
   - Por eso es importante verificar con: _, ok := m[clave]

6. CASOS DE USO:
   - Cache/Memoización
   - Contar ocurrencias
   - Búsqueda rápida (O(1) promedio)
   - Agrupación de datos
   - Configuraciones

EJECUTAR:
go run 06_maps.go
*/
