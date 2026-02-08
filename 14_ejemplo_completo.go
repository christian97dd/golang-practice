package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ========== EJEMPLO COMPLETO: SISTEMA DE PROCESAMIENTO DE PEDIDOS ==========
// Este ejemplo integra muchos conceptos de Go:
// - Structs y métodos
// - Interfaces
// - Goroutines y channels
// - Manejo de errores
// - Concurrencia con WaitGroup

// ========== TIPOS Y ESTRUCTURAS ==========

// Estado del pedido
type EstadoPedido int

const (
	Pendiente EstadoPedido = iota
	Procesando
	Completado
	Fallido
)

// String permite imprimir el estado de forma legible (implementa fmt.Stringer)
func (e EstadoPedido) String() string {
	return [...]string{"Pendiente", "Procesando", "Completado", "Fallido"}[e]
}

// Pedido representa un pedido de cliente
type Pedido struct {
	ID       int
	Cliente  string
	Producto string
	Cantidad int
	Estado   EstadoPedido
	Precio   float64
}

// Resultado del procesamiento
type Resultado struct {
	Pedido  *Pedido
	Exito   bool
	Mensaje string
}

// ========== INTERFAZ ==========

// Procesador define el contrato para procesar pedidos
type Procesador interface {
	Procesar(pedido *Pedido) error
}

// ========== IMPLEMENTACIONES ==========

// ProcesadorPago procesa el pago de un pedido
type ProcesadorPago struct {
	Nombre string
}

func (p *ProcesadorPago) Procesar(pedido *Pedido) error {
	fmt.Printf("  [Pago] Procesando pago para pedido #%d...\n", pedido.ID)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simular trabajo

	// Simular fallo ocasional
	if rand.Float32() < 0.1 { // 10% de probabilidad de fallo
		return fmt.Errorf("pago rechazado para pedido #%d", pedido.ID)
	}

	fmt.Printf("  [Pago] ✓ Pago completado para pedido #%d\n", pedido.ID)
	return nil
}

// ProcesadorInventario verifica y reserva inventario
type ProcesadorInventario struct {
	Stock map[string]int
	mutex sync.Mutex // Proteger acceso concurrente al mapa
}

func (p *ProcesadorInventario) Procesar(pedido *Pedido) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	fmt.Printf("  [Inventario] Verificando stock para pedido #%d...\n", pedido.ID)
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

	stockDisponible := p.Stock[pedido.Producto]
	if stockDisponible < pedido.Cantidad {
		return fmt.Errorf("stock insuficiente para %s (disponible: %d, solicitado: %d)",
			pedido.Producto, stockDisponible, pedido.Cantidad)
	}

	// Reservar stock
	p.Stock[pedido.Producto] -= pedido.Cantidad
	fmt.Printf("  [Inventario] ✓ Stock reservado para pedido #%d\n", pedido.ID)
	return nil
}

// ProcesadorEnvio prepara el envío del pedido
type ProcesadorEnvio struct{}

func (p *ProcesadorEnvio) Procesar(pedido *Pedido) error {
	fmt.Printf("  [Envío] Preparando envío para pedido #%d...\n", pedido.ID)
	time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)

	fmt.Printf("  [Envío] ✓ Pedido #%d listo para enviar\n", pedido.ID)
	return nil
}

// ========== SISTEMA DE PROCESAMIENTO ==========

// SistemaPedidos coordina el procesamiento de pedidos
type SistemaPedidos struct {
	procesadores []Procesador
	pedidos      chan *Pedido
	resultados   chan Resultado
	trabajadores int
}

// NewSistemaPedidos crea un nuevo sistema de pedidos
func NewSistemaPedidos(trabajadores int) *SistemaPedidos {
	return &SistemaPedidos{
		procesadores: []Procesador{
			&ProcesadorInventario{
				Stock: map[string]int{
					"Laptop":    10,
					"Mouse":     50,
					"Teclado":   30,
					"Monitor":   15,
					"Auricular": 25,
				},
			},
			&ProcesadorPago{Nombre: "Sistema de Pago"},
			&ProcesadorEnvio{},
		},
		pedidos:      make(chan *Pedido, 100),
		resultados:   make(chan Resultado, 100),
		trabajadores: trabajadores,
	}
}

// Iniciar lanza los trabajadores
func (s *SistemaPedidos) Iniciar() {
	var wg sync.WaitGroup

	// Lanzar trabajadores
	for i := 1; i <= s.trabajadores; i++ {
		wg.Add(1)
		go s.trabajador(i, &wg)
	}

	// Goroutine para cerrar resultados cuando todos los trabajadores terminen
	go func() {
		wg.Wait()
		close(s.resultados)
	}()
}

// Trabajador procesa pedidos
func (s *SistemaPedidos) trabajador(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for pedido := range s.pedidos {
		fmt.Printf("\n[Trabajador %d] Procesando pedido #%d de %s\n", id, pedido.ID, pedido.Cliente)
		pedido.Estado = Procesando

		// Procesar con cada procesador
		exito := true
		var mensaje string

		for _, procesador := range s.procesadores {
			if err := procesador.Procesar(pedido); err != nil {
				// Error en algún procesador
				exito = false
				mensaje = err.Error()
				pedido.Estado = Fallido
				fmt.Printf("[Trabajador %d] ✗ Error en pedido #%d: %s\n", id, pedido.ID, err)
				break
			}
		}

		if exito {
			pedido.Estado = Completado
			mensaje = "Pedido procesado exitosamente"
			fmt.Printf("[Trabajador %d] ✓ Pedido #%d completado\n", id, pedido.ID)
		}

		// Enviar resultado
		s.resultados <- Resultado{
			Pedido:  pedido,
			Exito:   exito,
			Mensaje: mensaje,
		}
	}
}

// AgregarPedido añade un pedido a la cola
func (s *SistemaPedidos) AgregarPedido(pedido *Pedido) {
	s.pedidos <- pedido
}

// Finalizar cierra el sistema
func (s *SistemaPedidos) Finalizar() {
	close(s.pedidos)
}

// Resultados retorna el channel de resultados
func (s *SistemaPedidos) Resultados() <-chan Resultado {
	return s.resultados
}

// ========== UTILIDADES ==========

// GenerarPedidos crea pedidos de ejemplo
func GenerarPedidos(cantidad int) []*Pedido {
	productos := []string{"Laptop", "Mouse", "Teclado", "Monitor", "Auricular"}
	nombres := []string{"Ana", "Carlos", "Luis", "María", "Pedro", "Laura"}
	pedidos := make([]*Pedido, cantidad)

	for i := 0; i < cantidad; i++ {
		pedidos[i] = &Pedido{
			ID:       i + 1,
			Cliente:  nombres[rand.Intn(len(nombres))],
			Producto: productos[rand.Intn(len(productos))],
			Cantidad: rand.Intn(3) + 1, // 1-3 unidades
			Estado:   Pendiente,
			Precio:   float64(rand.Intn(1000) + 100),
		}
	}

	return pedidos
}

// MostrarEstadisticas muestra el resumen de resultados
func MostrarEstadisticas(resultados []Resultado) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("RESUMEN DE PROCESAMIENTO")
	fmt.Println(strings.Repeat("=", 60))

	exitosos := 0
	fallidos := 0
	totalVentas := 0.0

	for _, res := range resultados {
		if res.Exito {
			exitosos++
			totalVentas += res.Pedido.Precio
		} else {
			fallidos++
		}
	}

	fmt.Printf("Total de pedidos:    %d\n", len(resultados))
	fmt.Printf("Exitosos:            %d (%.1f%%)\n", exitosos, float64(exitosos)/float64(len(resultados))*100)
	fmt.Printf("Fallidos:            %d (%.1f%%)\n", fallidos, float64(fallidos)/float64(len(resultados))*100)
	fmt.Printf("Total de ventas:     $%.2f\n", totalVentas)

	fmt.Println("\nDetalle de pedidos fallidos:")
	for _, res := range resultados {
		if !res.Exito {
			fmt.Printf("  - Pedido #%d: %s\n", res.Pedido.ID, res.Mensaje)
		}
	}
}

// ========== MAIN ==========

import "strings"

func main() {
	fmt.Println("========================================")
	fmt.Println("SISTEMA DE PROCESAMIENTO DE PEDIDOS")
	fmt.Println("========================================\n")

	// Seed para números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Crear sistema con 3 trabajadores
	sistema := NewSistemaPedidos(3)

	// Iniciar trabajadores
	sistema.Iniciar()

	// Generar pedidos
	pedidos := GenerarPedidos(15)

	fmt.Printf("Generados %d pedidos\n", len(pedidos))
	fmt.Println("\nIniciando procesamiento...\n")

	// Enviar pedidos al sistema
	go func() {
		for _, pedido := range pedidos {
			sistema.AgregarPedido(pedido)
			time.Sleep(100 * time.Millisecond) // Simular llegada gradual
		}
		sistema.Finalizar()
	}()

	// Recolectar resultados
	var resultados []Resultado
	for resultado := range sistema.Resultados() {
		resultados = append(resultados, resultado)
	}

	// Mostrar estadísticas
	MostrarEstadisticas(resultados)

	fmt.Println("\n¡Procesamiento completado!")
}

/*
========== CONCEPTOS DEMOSTRADOS ==========

1. STRUCTS:
   - Pedido, Resultado
   - Composición de datos

2. INTERFACES:
   - Procesador (polimorfismo)
   - fmt.Stringer (EstadoPedido)

3. MÉTODOS:
   - Procesadores implementan Procesar()
   - Métodos del SistemaPedidos

4. GOROUTINES:
   - Trabajadores concurrentes
   - Procesamiento paralelo

5. CHANNELS:
   - Canal de pedidos
   - Canal de resultados
   - Comunicación entre goroutines

6. SYNC:
   - WaitGroup para sincronización
   - Mutex para proteger mapa compartido

7. ERRORES:
   - Manejo de errores en procesamiento
   - Propagación de errores

8. PATRONES:
   - Worker Pool (pool de trabajadores)
   - Producer-Consumer (productor-consumidor)
   - Pipeline de procesamiento

FLUJO DEL PROGRAMA:

1. Se crea el sistema con N trabajadores
2. Se generan pedidos aleatorios
3. Los pedidos se envían al channel
4. Los trabajadores los toman del channel
5. Cada trabajador procesa con múltiples procesadores
6. Los resultados se envían a otro channel
7. Se recolectan y muestran estadísticas

MEJORAS POSIBLES:
- Agregar context para cancelación
- Persistencia de datos (base de datos)
- API REST para recibir pedidos
- Métricas y logging
- Tests unitarios
- Circuit breaker para fallos
- Retry logic

EJECUTAR:
go run 14_ejemplo_completo.go

OBSERVA:
- Cómo los trabajadores procesan en paralelo
- Manejo de errores (stock insuficiente, pagos fallidos)
- Sincronización con channels y WaitGroup
- Uso de mutex para proteger el mapa de inventario
*/
