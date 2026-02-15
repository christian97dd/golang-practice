# CRUD Básico en Go

Este es un ejemplo sencillo de un CRUD (Create, Read, Update, Delete) en Go para principiantes.

## ¿Qué es un CRUD?

CRUD es un acrónimo que representa las cuatro operaciones básicas de persistencia de datos:
- **C**reate (Crear)
- **R**ead (Leer)
- **U**pdate (Actualizar)
- **D**elete (Eliminar)

## Características

- ✅ Gestión de productos con ID, Nombre y Precio
- ✅ Almacenamiento en memoria (los datos se pierden al cerrar el programa)
- ✅ Menú interactivo fácil de usar
- ✅ Código comentado para facilitar el aprendizaje

## Cómo ejecutar

```bash
# Desde la carpeta crud:
go run main.go
```

## Estructura del código

### 1. Estructura de datos
```go
type Producto struct {
    ID     int
    Nombre string
    Precio float64
}
```

### 2. Operaciones CRUD

- **crearProducto()**: Añade un nuevo producto al slice
- **listarProductos()**: Muestra todos los productos
- **buscarProducto()**: Encuentra un producto por su ID
- **actualizarProducto()**: Modifica un producto existente
- **eliminarProducto()**: Elimina un producto del slice

### 3. Almacenamiento

Los productos se guardan en un **slice en memoria**:
```go
var productos []Producto
```

Esto significa que cuando cierres el programa, los datos se perderán.

## Conceptos de Go que aprenderás

1. **Structs**: Cómo definir estructuras de datos
2. **Slices**: Manipulación de arrays dinámicos
3. **Punteros**: Uso de `*Producto` para modificar datos
4. **Funciones**: Organización del código en funciones reutilizables
5. **Input/Output**: Lectura de datos del usuario
6. **Conversión de tipos**: `strconv.Atoi()`, `strconv.ParseFloat()`

## Ejemplo de uso

```
=== MENÚ ===
1. Crear producto
2. Listar productos
3. Buscar producto por ID
4. Actualizar producto
5. Eliminar producto
6. Salir

Selecciona una opción: 1
Nombre del producto: Laptop
Precio: 899.99
✓ Producto creado exitosamente con ID: 1
```

## Próximos pasos

Una vez que entiendas este CRUD básico, puedes:
- Añadir más campos a la estructura Producto
- Guardar los datos en un archivo JSON
- Conectar con una base de datos (MySQL, PostgreSQL)
- Crear una API REST con este CRUD
