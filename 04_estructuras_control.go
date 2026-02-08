package main

import "fmt"

func main() {
	fmt.Println("=== ESTRUCTURAS DE CONTROL ===\n")

	// ========== IF / ELSE ==========
	fmt.Println("--- IF / ELSE ---")

	edad := 18

	// If básico (sin paréntesis obligatorios)
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	}

	// If con else
	if edad < 18 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Eres mayor de edad")
	}

	// If con else if
	nota := 85
	if nota >= 90 {
		fmt.Println("Calificación: A")
	} else if nota >= 80 {
		fmt.Println("Calificación: B")
	} else if nota >= 70 {
		fmt.Println("Calificación: C")
	} else {
		fmt.Println("Calificación: F")
	}

	// If con inicialización (muy común en Go)
	// La variable solo existe dentro del bloque if
	if temperatura := 25; temperatura > 30 {
		fmt.Println("Hace calor")
	} else if temperatura > 20 {
		fmt.Println("Temperatura agradable:", temperatura)
	} else {
		fmt.Println("Hace frío")
	}

	// ========== FOR (ÚNICO BUCLE EN GO) ==========
	fmt.Println("\n--- FOR ---")

	// For clásico (como en C/Java)
	fmt.Print("Conteo: ")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// For como while
	contador := 0
	for contador < 3 {
		fmt.Println("Contador:", contador)
		contador++
	}

	// For infinito (como while(true))
	// Lo detenemos con break
	n := 0
	for {
		if n >= 3 {
			break // Sale del bucle
		}
		fmt.Println("Infinito:", n)
		n++
	}

	// Continue: salta a la siguiente iteración
	fmt.Print("Números impares: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Salta los números pares
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// ========== RANGE (ITERAR SOBRE COLECCIONES) ==========
	fmt.Println("\n--- RANGE ---")

	// Range sobre slice (array)
	numeros := []int{10, 20, 30, 40, 50}

	// Range retorna índice y valor
	fmt.Println("Con índice y valor:")
	for indice, valor := range numeros {
		fmt.Printf("  numeros[%d] = %d\n", indice, valor)
	}

	// Si no necesitas el índice, usa _ (blank identifier)
	fmt.Print("Solo valores: ")
	for _, valor := range numeros {
		fmt.Print(valor, " ")
	}
	fmt.Println()

	// Range sobre string (itera sobre caracteres)
	texto := "Hola"
	for indice, char := range texto {
		fmt.Printf("  texto[%d] = %c\n", indice, char)
	}

	// Range sobre map (diccionario)
	edades := map[string]int{
		"Ana":    25,
		"Carlos": 30,
		"Luis":   28,
	}
	fmt.Println("Iterando sobre map:")
	for nombre, edad := range edades {
		fmt.Printf("  %s tiene %d años\n", nombre, edad)
	}

	// ========== SWITCH ==========
	fmt.Println("\n--- SWITCH ---")

	dia := "martes"

	// Switch básico
	switch dia {
	case "lunes":
		fmt.Println("Inicio de semana")
	case "martes", "miércoles", "jueves": // Múltiples casos
		fmt.Println("Mitad de semana")
	case "viernes":
		fmt.Println("¡Fin de semana cerca!")
	case "sábado", "domingo":
		fmt.Println("Fin de semana")
	default:
		fmt.Println("Día no reconocido")
	}

	// Switch sin condición (como if-else-if)
	hora := 14
	switch {
	case hora < 12:
		fmt.Println("Buenos días")
	case hora < 18:
		fmt.Println("Buenas tardes")
	default:
		fmt.Println("Buenas noches")
	}

	// Switch con inicialización
	switch num := 10; {
	case num < 0:
		fmt.Println("Negativo")
	case num == 0:
		fmt.Println("Cero")
	default:
		fmt.Println("Positivo:", num)
	}

	// Switch con tipos (type switch) - avanzado
	var valor interface{} = "hola"
	switch v := valor.(type) {
	case string:
		fmt.Printf("Es un string: %s\n", v)
	case int:
		fmt.Printf("Es un int: %d\n", v)
	default:
		fmt.Printf("Tipo desconocido: %T\n", v)
	}

	// ========== DEFER ==========
	fmt.Println("\n--- DEFER ---")
	// defer pospone la ejecución hasta que la función termine
	// Útil para limpieza (cerrar archivos, conexiones, etc.)

	defer fmt.Println("Esto se ejecuta AL FINAL")
	fmt.Println("Esto se ejecuta primero")
	fmt.Println("Esto se ejecuta segundo")
	// Cuando main termina, se ejecuta el defer
}

/*
RESUMEN:

1. IF:
   - Sin paréntesis obligatorios
   - Puede incluir inicialización: if x := valor; condicion {}

2. FOR:
   - Único bucle en Go (no hay while ni do-while)
   - For clásico: for i := 0; i < 10; i++ {}
   - Como while: for condicion {}
   - Infinito: for {}
   - break: sale del bucle
   - continue: siguiente iteración

3. RANGE:
   - Itera sobre slices, arrays, maps, strings
   - Retorna (índice, valor) o (clave, valor) en maps
   - Usa _ para ignorar valores

4. SWITCH:
   - No necesita break (no hay fall-through por defecto)
   - Puede tener múltiples casos en una línea
   - Puede ser sin condición (como if-else-if)
   - Type switch para verificar tipos

5. DEFER:
   - Pospone ejecución hasta el final de la función
   - Útil para limpieza y cierre de recursos

EJECUTAR:
go run 04_estructuras_control.go
*/
