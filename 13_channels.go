package main

import (
	"fmt"
	"time"
)

// ========== CHANNELS ==========
// Los channels permiten que las goroutines se COMUNIQUEN entre sí
// Son como tuberías (pipes) por donde fluyen datos
// Lema de Go: "No compartas memoria por comunicación, comunica por compartir memoria"

func main() {
	fmt.Println("=== CHANNELS (Canales) ===\n")

	// ========== CHANNEL BÁSICO ==========
	fmt.Println("--- CHANNEL BÁSICO ---")

	// Crear channel de enteros
	// make(chan tipo)
	ch := make(chan int)

	// Enviar valor a channel (en goroutine, sino bloquea)
	go func() {
		ch <- 42 // <- es el operador de envío
		fmt.Println("Valor enviado al channel")
	}()

	// Recibir valor del channel
	valor := <-ch // <- es el operador de recepción
	fmt.Printf("Valor recibido: %d\n\n", valor)

	// ========== SINCRONIZACIÓN CON CHANNELS ==========
	fmt.Println("--- SINCRONIZACIÓN ---")

	// Los channels sincronizan goroutines automáticamente
	hecho := make(chan bool)

	go func() {
		fmt.Println("Trabajando...")
		time.Sleep(1 * time.Second)
		fmt.Println("Trabajo terminado")
		hecho <- true // Señal de que terminó
	}()

	fmt.Println("Esperando...")
	<-hecho // Espera hasta recibir la señal
	fmt.Println("Trabajo completado\n")

	// ========== MÚLTIPLES VALORES ==========
	fmt.Println("--- MÚLTIPLES VALORES ---")

	numeros := make(chan int)

	// Productor: envía números
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Enviando: %d\n", i)
			numeros <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(numeros) // IMPORTANTE: cerrar el channel cuando termines
	}()

	// Consumidor: recibe números
	for num := range numeros { // range itera hasta que el channel se cierre
		fmt.Printf("  Recibido: %d\n", num)
	}
	fmt.Println()

	// ========== BUFFERED CHANNELS ==========
	fmt.Println("--- BUFFERED CHANNELS ---")

	// Channel sin buffer: envío bloquea hasta que alguien reciba
	// Channel con buffer: envío no bloquea hasta que el buffer esté lleno

	// Channel con buffer de 3 elementos
	buffered := make(chan string, 3)

	// Podemos enviar 3 valores sin bloquear (sin goroutine)
	buffered <- "primero"
	buffered <- "segundo"
	buffered <- "tercero"
	// buffered <- "cuarto" // Esto bloquearía porque el buffer está lleno

	fmt.Println(<-buffered) // primero
	fmt.Println(<-buffered) // segundo
	fmt.Println(<-buffered) // tercero
	fmt.Println()

	// ========== SELECT ==========
	fmt.Println("--- SELECT (multiplexar channels) ---")

	// select permite esperar en múltiples channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "mensaje de ch1"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch2 <- "mensaje de ch2"
	}()

	// select escucha múltiples channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recibido:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recibido:", msg2)
		}
	}
	fmt.Println()

	// ========== SELECT CON DEFAULT ==========
	fmt.Println("--- SELECT CON DEFAULT ---")

	// default hace que select no bloquee
	mensajes := make(chan string, 1)
	mensajes <- "hola"

	select {
	case msg := <-mensajes:
		fmt.Println("Mensaje:", msg)
	default:
		fmt.Println("No hay mensajes")
	}

	// Sin mensajes
	select {
	case msg := <-mensajes:
		fmt.Println("Mensaje:", msg)
	default:
		fmt.Println("No hay mensajes (default)")
	}
	fmt.Println()

	// ========== SELECT CON TIMEOUT ==========
	fmt.Println("--- SELECT CON TIMEOUT ---")

	resultado := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Tarda mucho
		resultado <- "resultado"
	}()

	select {
	case res := <-resultado:
		fmt.Println("Recibido:", res)
	case <-time.After(1 * time.Second): // Timeout de 1 segundo
		fmt.Println("Timeout: operación demoró mucho")
	}
	fmt.Println()

	// ========== PATRÓN PRODUCTOR-CONSUMIDOR ==========
	fmt.Println("--- PATRÓN PRODUCTOR-CONSUMIDOR ---")

	trabajos := make(chan int, 5)
	resultados := make(chan int, 5)

	// Trabajadores (consumidores)
	for w := 1; w <= 3; w++ {
		go trabajador(w, trabajos, resultados)
	}

	// Enviar trabajos
	for j := 1; j <= 5; j++ {
		trabajos <- j
	}
	close(trabajos)

	// Recibir resultados
	for r := 1; r <= 5; r++ {
		<-resultados
	}
	fmt.Println()

	// ========== DIRECCIONALIDAD DE CHANNELS ==========
	fmt.Println("--- DIRECCIONALIDAD ---")

	// Channels unidireccionales en firmas de funciones
	// chan<- solo envío
	// <-chan solo recepción

	ping := make(chan string, 1)
	pong := make(chan string, 1)

	enviarPing(ping)
	recibirYReenviar(ping, pong)
	recibirPong(pong)
	fmt.Println()

	// ========== CERRAR CHANNELS ==========
	fmt.Println("--- CERRAR CHANNELS ---")

	datos := make(chan int, 3)

	datos <- 1
	datos <- 2
	datos <- 3
	close(datos) // Cerrar channel

	// Recibir hasta que esté cerrado
	for valor := range datos {
		fmt.Println("Dato:", valor)
	}

	// Verificar si un channel está cerrado
	_, abierto := <-datos
	fmt.Printf("Channel abierto: %t\n", abierto)
	fmt.Println()

	// ========== EJEMPLO PRÁCTICO ==========
	fmt.Println("--- EJEMPLO PRÁCTICO: PIPELINE ---")

	pipeline()
}

// ========== FUNCIONES DE EJEMPLO ==========

func trabajador(id int, trabajos <-chan int, resultados chan<- int) {
	for trabajo := range trabajos {
		fmt.Printf("  Trabajador %d procesando trabajo %d\n", id, trabajo)
		time.Sleep(500 * time.Millisecond)
		resultados <- trabajo * 2
		fmt.Printf("  Trabajador %d terminó trabajo %d\n", id, trabajo)
	}
}

// Channel de solo envío (chan<- tipo)
func enviarPing(ping chan<- string) {
	ping <- "ping"
	fmt.Println("Enviado: ping")
}

// Recibe de uno, envía a otro
func recibirYReenviar(ping <-chan string, pong chan<- string) {
	msg := <-ping
	fmt.Println("Recibido:", msg)
	pong <- "pong"
	fmt.Println("Enviado: pong")
}

// Channel de solo recepción (<-chan tipo)
func recibirPong(pong <-chan string) {
	msg := <-pong
	fmt.Println("Recibido:", msg)
}

// ========== PIPELINE ==========

func pipeline() {
	// Pipeline de procesamiento: numeros -> cuadrados -> sumas

	// Stage 1: Generar números
	numeros := generarNumeros(1, 2, 3, 4, 5)

	// Stage 2: Calcular cuadrados
	cuadrados := calcularCuadrados(numeros)

	// Stage 3: Sumar
	suma := sumarTodo(cuadrados)

	// Resultado
	fmt.Printf("Suma de cuadrados: %d\n", suma)
}

func generarNumeros(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func calcularCuadrados(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func sumarTodo(in <-chan int) int {
	suma := 0
	for n := range in {
		suma += n
	}
	return suma
}

/*
CONCEPTOS CLAVE SOBRE CHANNELS:

1. CREAR CHANNELS:
   ch := make(chan tipo)           // Sin buffer
   ch := make(chan tipo, capacidad) // Con buffer

2. OPERADORES:
   ch <- valor    // Enviar a channel
   valor := <-ch  // Recibir de channel
   close(ch)      // Cerrar channel

3. TIPOS:
   SIN BUFFER (unbuffered):
   - Envío bloquea hasta que alguien recibe
   - Recepción bloquea hasta que alguien envía
   - Sincronización garantizada

   CON BUFFER (buffered):
   - Envío no bloquea si hay espacio en buffer
   - Recepción no bloquea si hay datos en buffer
   - Útil para desacoplar productor/consumidor

4. CERRAR CHANNELS:
   close(ch)
   - Solo el emisor debe cerrar
   - Recibir de channel cerrado retorna valor cero
   - Enviar a channel cerrado causa PANIC
   - for range termina cuando el channel se cierra

5. SELECT:
   - Espera en múltiples channels
   - Ejecuta el primer case que esté listo
   - default: no bloquea si ningún case está listo
   - time.After: timeout

6. DIRECCIONALIDAD:
   chan<- T   // Solo envío
   <-chan T   // Solo recepción
   - Útil en firmas de funciones
   - Previene mal uso

7. PATRONES COMUNES:
   - Productor-Consumidor
   - Pipeline (cadena de procesamiento)
   - Fan-out (múltiples consumidores)
   - Fan-in (múltiples productores)
   - Cancelación con channels

8. BUENAS PRÁCTICAS:
   ✓ Solo el emisor cierra el channel
   ✓ Cerrar channels cuando no se usarán más
   ✓ Usar select para múltiples channels
   ✓ Usar buffer para evitar bloqueos innecesarios
   ✓ Channels de solo lectura/escritura en funciones

9. PROBLEMAS COMUNES:
   ✗ Deadlock: todos bloqueados esperando
   ✗ Goroutine leak: goroutines esperando forever
   ✗ Cerrar channel ya cerrado: panic
   ✗ Enviar a channel cerrado: panic

10. COMPARACIÓN:
    JavaScript: Promises, async/await
    Python: Queue, asyncio
    Java: BlockingQueue
    Go: Channels (más integrado al lenguaje)

FILOSOFÍA DE GO:
"No compartas memoria por comunicación;
 en su lugar, comunica por compartir memoria"

Los channels implementan esta filosofía:
- En vez de que múltiples goroutines accedan la misma variable (peligroso)
- Las goroutines se comunican pasando datos por channels (seguro)

EJECUTAR:
go run 13_channels.go
*/
