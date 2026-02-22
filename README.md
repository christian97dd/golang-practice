# Aprendizaje de Go (Golang)

Este repositorio contiene material de aprendizaje de Go organizado en carpetas.

## 📁 Estructura del Proyecto

```
📦 go-aprendizaje/
├── 📂 ejercicios/        # Ejercicios teóricos numerados (01-14)
├── 📂 ejercitación/      # Ejercicios prácticos propios
├── 📂 crud/              # CRUD básico con menú interactivo
├── 📂 crud-api/          # CRUD API REST con Gin
├── 📂 guess-number-game/ # Juego de adivinar un número
└── 📂 todo-list/         # Lista de tareas con CRUD en terminal
```

## Requisitos

1. **Instalar Go**: Descarga desde [https://go.dev/dl/](https://go.dev/dl/)
2. **Verificar instalación**: `go version`

## 📚 Ejercicios Básicos

La carpeta `ejercicios/` contiene 14 archivos para aprender Go paso a paso.

### Cómo ejecutar los ejercicios

```bash
# Navega a la carpeta de ejercicios
cd ejercicios

# Ejecuta cualquier archivo
go run 01_hola_mundo.go
```

### Orden recomendado de lectura

1. **01_hola_mundo.go** - Estructura básica de un programa Go
2. **02_variables_tipos.go** - Variables y tipos de datos
3. **03_funciones.go** - Declaración y uso de funciones
4. **04_estructuras_control.go** - If, for, switch
5. **05_arrays_slices.go** - Arreglos y slices (listas dinámicas)
6. **06_maps.go** - Mapas (diccionarios/objetos)
7. **07_structs.go** - Estructuras (similar a clases)
8. **08_metodos.go** - Métodos en structs
9. **09_interfaces.go** - Interfaces y polimorfismo
10. **10_punteros.go** - Punteros y referencias
11. **11_errores.go** - Manejo de errores en Go
12. **12_goroutines.go** - Concurrencia con goroutines
13. **13_channels.go** - Comunicación entre goroutines
14. **14_ejemplo_completo.go** - Aplicación práctica completa

## ✏️ Ejercitación

La carpeta `ejercitación/` contiene ejercicios prácticos propios para reforzar lo aprendido.

```bash
cd ejercitación
go run hola.go
```

| Archivo | Contenido |
|---------|-----------|
| `hola.go` | Variables, constantes, `iota` (días de la semana) |
| `tipos-datos.go` | Tipos de datos, bytes, runes, conversión de tipos, input del usuario |
| `ejercicio.go` | Ejercicio: hipotenusa, área y perímetro de un triángulo rectángulo con `math` |
| `if-switch.go` | Condicionales `if/else` y `switch`, uso de `time` y `runtime` |
| `for.go` | Bucle `for` con `continue` |
| `funciones.go` | Funciones, múltiples retornos, named returns |
| `matrices.go` | Arrays, matrices 2D, slices: `append`, `copy`, eliminación por índice, `make` |
| `mapas.go` | Mapas: creación, acceso, patrón `ok`, `delete`, iteración con `range` |
| `estructuras.go` | Structs: declaración, inicialización con campos nombrados y posicional |
| `punteros.go` | Punteros, referencias con `&` y `*`, métodos con receptor puntero |

## 🎮 Proyectos

### 1. CRUD con Menú Interactivo (`crud/`)

CRUD básico de productos que funciona en memoria con menú en terminal.

**Ejecutar:**
```bash
cd crud
go run main.go
```

📖 Ver más: [crud/README.md](crud/README.md)

### 2. CRUD API REST con Gin (`crud-api/`)

API REST completa para gestionar productos usando el framework Gin.

**Ejecutar:**
```bash
cd crud-api
go mod download  # Primera vez
go run main.go
```

**Probar:**
```bash
# Crear un producto
curl -X POST http://localhost:8080/productos \
  -H "Content-Type: application/json" \
  -d '{"nombre": "Laptop", "precio": 899.99}'

# Listar productos
curl http://localhost:8080/productos
```

📖 Ver más: [crud-api/README.md](crud-api/README.md)

### 3. Juego: Adivina el Número (`guess-number-game/`)

Juego interactivo en terminal para adivinar un número aleatorio entre 0 y 99.

**Ejecutar:**
```bash
cd guess-number-game
go run main.go
```

- 10 intentos máximos
- Pistas de "mayor" o "menor" en cada intento
- Opción de jugar de nuevo al terminar

📖 Ver más: [guess-number-game/README.md](guess-number-game/README.md)

### 4. Lista de Tareas (`todo-list/`)

CRUD de tareas interactivo en terminal. Permite agregar, completar, editar y eliminar tareas usando structs y métodos.

**Ejecutar:**
```bash
cd todo-list
go run main.go
```

- Agregar tareas con nombre y descripción
- Marcar tareas como completadas
- Editar y eliminar tareas por índice
- Lista las tareas después de cada operación

📖 Ver más: [todo-list/README.md](todo-list/README.md)

## Características de Go

- **Compilado**: Se compila a binario nativo (muy rápido)
- **Tipado estático**: Los tipos se verifican en compilación
- **Garbage collection**: Manejo automático de memoria
- **Concurrencia nativa**: Goroutines y channels integrados
- **Simplicidad**: Sintaxis minimalista y clara
