package main

import "fmt"

func main() {
	fmt.Println("=== ARRAYS Y SLICES ===\n")

	// ========== ARRAYS ==========
	// Arrays tienen TAMAÑO FIJO
	fmt.Println("--- ARRAYS (tamaño fijo) ---")

	// Declarar array de 5 enteros
	var numeros [5]int
	fmt.Println("Array vacío:", numeros) // [0 0 0 0 0]

	// Asignar valores
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println("Array modificado:", numeros)

	// Declarar e inicializar
	frutas := [3]string{"manzana", "pera", "uva"}
	fmt.Println("Frutas:", frutas)

	// Tamaño automático con ...
	colores := [...]string{"rojo", "verde", "azul", "amarillo"}
	fmt.Println("Colores:", colores)
	fmt.Println("Longitud:", len(colores)) // len() da la longitud

	// Acceder a elementos
	fmt.Printf("Primera fruta: %s\n", frutas[0])
	fmt.Printf("Última fruta: %s\n", frutas[len(frutas)-1])

	// Iterar sobre array
	fmt.Print("Iterando colores: ")
	for i := 0; i < len(colores); i++ {
		fmt.Print(colores[i], " ")
	}
	fmt.Println()

	// ========== SLICES ==========
	// Slices son DINÁMICOS (tamaño variable)
	// Son la forma más común de trabajar con listas en Go
	fmt.Println("\n--- SLICES (tamaño dinámico) ---")

	// Declarar slice vacío
	var nombres []string
	fmt.Println("Slice vacío:", nombres)
	fmt.Println("¿Es nil?:", nombres == nil) // true
	fmt.Println("Longitud:", len(nombres))   // 0

	// Crear slice con make
	// make([]tipo, longitud, capacidad)
	edades := make([]int, 3) // Slice de 3 enteros
	fmt.Println("Slice con make:", edades) // [0 0 0]

	// Slice literal (más común)
	ciudades := []string{"Madrid", "Barcelona", "Valencia"}
	fmt.Println("Ciudades:", ciudades)

	// ========== APPEND (agregar elementos) ==========
	fmt.Println("\n--- APPEND ---")

	var lista []int
	fmt.Println("Lista inicial:", lista)

	// append añade elementos y retorna un NUEVO slice
	lista = append(lista, 10)
	lista = append(lista, 20)
	lista = append(lista, 30)
	fmt.Println("Después de append:", lista)

	// Agregar múltiples elementos
	lista = append(lista, 40, 50, 60)
	fmt.Println("Múltiples append:", lista)

	// Concatenar slices
	otros := []int{70, 80}
	lista = append(lista, otros...) // ... desempaqueta el slice
	fmt.Println("Concatenado:", lista)

	// ========== SLICING (subslices) ==========
	fmt.Println("\n--- SLICING ---")

	numeros2 := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Original:", numeros2)

	// slice[inicio:fin] (fin NO incluido)
	sub1 := numeros2[1:4] // Elementos 1, 2, 3
	fmt.Println("numeros2[1:4]:", sub1) // [20 30 40]

	// Desde el inicio
	sub2 := numeros2[:3] // Primeros 3 elementos
	fmt.Println("numeros2[:3]:", sub2) // [10 20 30]

	// Hasta el final
	sub3 := numeros2[3:] // Desde el índice 3 hasta el final
	fmt.Println("numeros2[3:]:", sub3) // [40 50 60]

	// Todo el slice
	copia := numeros2[:] // Copia superficial
	fmt.Println("numeros2[:]:", copia)

	// ========== COPY ==========
	fmt.Println("\n--- COPY ---")

	original := []int{1, 2, 3}
	copia2 := make([]int, len(original))
	copy(copia2, original) // Copia los elementos

	fmt.Println("Original:", original)
	fmt.Println("Copia:", copia2)

	// Modificar copia no afecta original
	copia2[0] = 999
	fmt.Println("Original después:", original) // [1 2 3]
	fmt.Println("Copia después:", copia2)      // [999 2 3]

	// ========== CAPACIDAD vs LONGITUD ==========
	fmt.Println("\n--- CAPACIDAD vs LONGITUD ---")

	// len(): número de elementos
	// cap(): capacidad asignada (espacio reservado)

	s := make([]int, 3, 5) // longitud 3, capacidad 5
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)

	s = append(s, 10)
	fmt.Printf("Después de append: len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)

	// ========== SLICE MULTIDIMENSIONAL ==========
	fmt.Println("\n--- SLICES MULTIDIMENSIONALES ---")

	// Matriz (slice de slices)
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matriz:")
	for i, fila := range matriz {
		fmt.Printf("  Fila %d: %v\n", i, fila)
	}

	// Acceder a elemento
	fmt.Printf("Elemento [1][2]: %d\n", matriz[1][2]) // 6

	// ========== FUNCIONES CON SLICES ==========
	fmt.Println("\n--- FUNCIONES CON SLICES ---")

	valores := []int{5, 2, 8, 1, 9}
	fmt.Println("Valores:", valores)
	fmt.Println("Suma:", sumarSlice(valores))
	fmt.Println("Máximo:", encontrarMaximo(valores))

	// Modificar slice en función
	duplicar(valores)
	fmt.Println("Duplicados:", valores)
}

// Función que recibe un slice
func sumarSlice(numeros []int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func encontrarMaximo(numeros []int) int {
	if len(numeros) == 0 {
		return 0
	}
	max := numeros[0]
	for _, num := range numeros {
		if num > max {
			max = num
		}
	}
	return max
}

// Los slices se pasan por referencia
// Modificar el slice dentro de la función afecta el original
func duplicar(numeros []int) {
	for i := range numeros {
		numeros[i] *= 2
	}
}

/*
DIFERENCIAS CLAVE: ARRAYS vs SLICES

ARRAYS:
- Tamaño FIJO en la declaración
- Sintaxis: var arr [5]int
- Menos flexibles
- Se pasan por VALOR a funciones (se copian)

SLICES:
- Tamaño DINÁMICO
- Sintaxis: var slice []int
- Más comunes y útiles
- Se pasan por REFERENCIA a funciones
- Internamente son referencias a arrays
- Operaciones: append, copy, slicing

CONCEPTOS IMPORTANTES:

1. len(): Longitud actual (número de elementos)
2. cap(): Capacidad (espacio reservado)
3. append(): Añade elementos (retorna nuevo slice)
4. copy(): Copia elementos de un slice a otro
5. slice[inicio:fin]: Crea subslice (fin NO incluido)
6. make([]T, len, cap): Crea slice con capacidad específica

CUÁNDO USAR:
- Arrays: Tamaño fijo conocido, optimización de memoria
- Slices: 99% de los casos (listas dinámicas)

EJECUTAR:
go run 05_arrays_slices.go
*/
