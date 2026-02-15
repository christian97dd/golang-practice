# Aprendizaje de Go (Golang)

Este repositorio contiene material de aprendizaje de Go organizado en carpetas.

## 📁 Estructura del Proyecto

```
📦 go-aprendizaje/
├── 📂 ejercicios/     # Ejercicios básicos numerados
└── 📂 crud/           # Proyecto CRUD básico
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

## 🚀 Proyecto CRUD

La carpeta `crud/` contiene un proyecto práctico completo.

### ¿Qué es?

Un CRUD (Create, Read, Update, Delete) básico de productos que funciona en memoria. Perfecto para entender cómo estructurar una aplicación real en Go.

### Cómo ejecutar el CRUD

```bash
# Navega a la carpeta crud
cd crud

# Ejecuta el programa
go run main.go
```

Ver más detalles en [crud/README.md](crud/README.md)

## Características de Go

- **Compilado**: Se compila a binario nativo (muy rápido)
- **Tipado estático**: Los tipos se verifican en compilación
- **Garbage collection**: Manejo automático de memoria
- **Concurrencia nativa**: Goroutines y channels integrados
- **Simplicidad**: Sintaxis minimalista y clara
