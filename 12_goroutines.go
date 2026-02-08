package main

import (
	"fmt"
	"time"
)

// ========== GOROUTINES ==========
// Las goroutines son funciones que se ejecutan CONCURRENTEMENTE
// Son extremadamente ligeras (miles pueden ejecutarse simultáneamente)
// La palabra clave "go" lanza una goroutine

func main() {
	fmt.Println("=== GOROUTINES (Concurrencia) ===\n")

	// ========== GOROUTINE BÁSICA ==========
	fmt.Println("--- GOROUTINE BÁSICA ---")

	// Función normal (síncrona)
	fmt.Println("Inicio del programa")
	saludar("mundo") // Se ejecuta y espera
	fmt.Println("Después de saludar\n")

	// Con goroutine (asíncrona)
	fmt.Println("Usando goroutine:")
	go saludar("goroutine") // Se ejecuta en paralelo, no espera
	fmt.Println("Después de lanzar goroutine")

	// PROBLEMA: El programa puede terminar antes de que la goroutine termine
	time.Sleep(1 * time.Second) // Esperar para ver la salida
	fmt.Println()

	// ========== MÚLTIPLES GOROUTINES ==========
	fmt.Println("--- MÚLTIPLES GOROUTINES ---")

	// Lanzar varias goroutines
	for i := 1; i <= 5; i++ {
		go contar(i)
	}

	// Esperar para que las goroutines terminen
	time.Sleep(2 * time.Second)
	fmt.Println()

	// ========== GOROUTINES CON FUNCIONES ANÓNIMAS ==========
	fmt.Println("--- GOROUTINES ANÓNIMAS ---")

	// Función anónima como goroutine
	go func() {
		fmt.Println("Goroutine anónima ejecutándose")
	}()

	// Con parámetros
	mensaje := "Hola desde goroutine"
	go func(msg string) {
		fmt.Println(msg)
	}(mensaje) // Pasar parámetro

	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// ========== PROBLEMA: COMPARTIR ESTADO ==========
	fmt.Println("--- COMPARTIR ESTADO (CUIDADO) ---")

	// PROBLEMA: Múltiples goroutines accediendo la misma variable
	contador := 0

	// Lanzar 100 goroutines que incrementan el contador
	for i := 0; i < 100; i++ {
		go func() {
			contador++ // Race condition: no es seguro
		}()
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Contador (puede ser incorrecto): %d\n", contador)
	fmt.Println("Debería ser 100, pero probablemente no lo es")
	fmt.Println("Solución: usar channels o sync.Mutex (ver ejemplos siguientes)\n")

	// ========== SINCRONIZACIÓN CON WAITGROUP ==========
	fmt.Println("--- SINCRONIZACIÓN CON sync.WaitGroup ---")

	// WaitGroup permite esperar a que goroutines terminen
	// (Importar "sync" - ver ejemplo completo abajo)

	ejemploWaitGroup()
	fmt.Println()

	// ========== GOROUTINE INFINITA ==========
	fmt.Println("--- GOROUTINE INFINITA ---")

	// Goroutine que se ejecuta continuamente
	go func() {
		for {
			fmt.Print(".")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	fmt.Println("Goroutine infinita ejecutándose...")
	time.Sleep(2 * time.Second)
	fmt.Println("\nSaliendo del programa (goroutine se detiene)")
	fmt.Println()

	// ========== SCHEDULER DE GO ==========
	fmt.Println("--- SCHEDULER DE GO ---")

	// El scheduler de Go decide cuándo ejecutar cada goroutine
	// Usa todos los núcleos del CPU disponibles

	for i := 1; i <= 3; i++ {
		go func(id int) {
			for j := 1; j <= 3; j++ {
				fmt.Printf("Goroutine %d: paso %d\n", id, j)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println()

	// ========== EJEMPLO PRÁCTICO ==========
	fmt.Println("--- EJEMPLO PRÁCTICO: PROCESAMIENTO PARALELO ---")

	procesarEnParalelo()
}

// ========== FUNCIONES DE EJEMPLO ==========

func saludar(nombre string) {
	fmt.Printf("Hola, %s!\n", nombre)
}

func contar(id int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

// ========== EJEMPLO CON WAITGROUP ==========

import "sync"

func ejemploWaitGroup() {
	var wg sync.WaitGroup // Contador de goroutines

	// Lanzar 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Incrementar contador

		go func(id int) {
			defer wg.Done() // Decrementar cuando termine
			fmt.Printf("Trabajador %d iniciado\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Trabajador %d terminado\n", id)
		}(i)
	}

	wg.Wait() // Esperar a que todas terminen
	fmt.Println("Todos los trabajadores terminaron")
}

// ========== EJEMPLO PRÁCTICO ==========

func procesarEnParalelo() {
	tareas := []string{"tarea1", "tarea2", "tarea3", "tarea4"}

	var wg sync.WaitGroup

	fmt.Println("Procesando tareas en paralelo...")
	inicio := time.Now()

	for _, tarea := range tareas {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			procesarTarea(t)
		}(tarea)
	}

	wg.Wait()

	duracion := time.Since(inicio)
	fmt.Printf("Todas las tareas completadas en %v\n", duracion)
}

func procesarTarea(nombre string) {
	fmt.Printf("  Procesando %s...\n", nombre)
	time.Sleep(500 * time.Millisecond) // Simular trabajo
	fmt.Printf("  %s completada\n", nombre)
}

/*
CONCEPTOS CLAVE SOBRE GOROUTINES:

1. QUÉ SON:
   - Funciones que se ejecutan concurrentemente
   - Extremadamente ligeras (~2KB cada una)
   - Administradas por el runtime de Go
   - NO son hilos del sistema operativo

2. SINTAXIS:
   go funcion()              // Lanzar goroutine
   go func() { ... }()       // Goroutine anónima

3. CARACTERÍSTICAS:
   - Se ejecutan en paralelo con otras goroutines
   - El scheduler decide cuándo ejecutar cada una
   - Usan todos los núcleos del CPU
   - Pueden ser miles ejecutándose simultáneamente

4. COMUNICACIÓN:
   - NO compartir estado directamente (race conditions)
   - Usar channels para comunicación (ver 13_channels.go)
   - O usar sync.Mutex para proteger datos compartidos

5. SINCRONIZACIÓN:
   sync.WaitGroup:
   - wg.Add(1): Incrementar contador
   - wg.Done(): Decrementar contador
   - wg.Wait(): Esperar a que contador sea 0

6. PROBLEMAS COMUNES:
   ✗ Race conditions: múltiples goroutines modificando mismo dato
   ✗ Programa termina antes que goroutines
   ✗ Goroutines bloqueadas (deadlock)
   ✗ Goroutines que nunca terminan (leak)

7. BUENAS PRÁCTICAS:
   ✓ Usar channels para comunicación
   ✓ Usar WaitGroup para esperar
   ✓ No compartir memoria, compartir por comunicación
   ✓ Siempre pensar en cuándo terminará la goroutine
   ✓ Usar context para cancelación

8. DIFERENCIAS CON OTROS LENGUAJES:
   - JavaScript: async/await, Promises
   - Python: asyncio, threading
   - Java: Threads
   - Go: Goroutines (más ligeras y fáciles)

9. CUÁNDO USAR:
   ✓ Operaciones I/O (red, archivos, DB)
   ✓ Procesamiento paralelo
   ✓ Servidores concurrentes
   ✓ Tareas en background

10. SCHEDULER:
    - Go usa M:N scheduling
    - Múltiples goroutines en múltiples hilos OS
    - GOMAXPROCS: número de núcleos a usar

CONCEPTOS IMPORTANTES:

- Concurrencia: Composición de procesos independientes
- Paralelismo: Ejecución simultánea de múltiples tareas
- Go permite escribir código concurrente fácilmente
- El runtime maneja la complejidad del paralelismo

PRÓXIMO PASO:
Ver 13_channels.go para aprender a comunicar goroutines de forma segura

EJECUTAR:
go run 12_goroutines.go
*/
