# CRUD API REST con Gin

API REST completa para gestionar productos usando Go y el framework Gin.

## 📁 Estructura del Proyecto

```
crud-api/
├── main.go              # Punto de entrada de la aplicación
├── go.mod               # Dependencias del proyecto
├── models/
│   └── producto.go      # Modelo de datos
├── handlers/
│   └── productos.go     # Lógica de negocio (CRUD)
└── routes/
    └── routes.go        # Definición de rutas HTTP
```

## 🚀 Instalación y Ejecución

### 1. Instalar dependencias

```bash
cd crud-api
go mod download
```

### 2. Ejecutar el servidor

```bash
go run main.go
```

El servidor se iniciará en `http://localhost:8080`

## 📡 Endpoints Disponibles

| Método | Endpoint           | Descripción                    |
|--------|-------------------|--------------------------------|
| GET    | `/`               | Información de la API          |
| GET    | `/productos`      | Listar todos los productos     |
| GET    | `/productos/:id`  | Obtener un producto por ID     |
| POST   | `/productos`      | Crear un nuevo producto        |
| PUT    | `/productos/:id`  | Actualizar un producto         |
| DELETE | `/productos/:id`  | Eliminar un producto           |

## 🧪 Ejemplos de Uso

### 1️⃣ Obtener información de la API

```bash
curl http://localhost:8080
```

**Respuesta:**
```json
{
  "mensaje": "¡Bienvenido al CRUD API de Productos!",
  "versión": "1.0",
  "endpoints": { ... }
}
```

### 2️⃣ Crear un producto (POST)

```bash
curl -X POST http://localhost:8080/productos \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Laptop",
    "precio": 899.99
  }'
```

**Respuesta:**
```json
{
  "id": 1,
  "nombre": "Laptop",
  "precio": 899.99
}
```

### 3️⃣ Listar todos los productos (GET)

```bash
curl http://localhost:8080/productos
```

**Respuesta:**
```json
{
  "productos": [
    {
      "id": 1,
      "nombre": "Laptop",
      "precio": 899.99
    }
  ],
  "total": 1
}
```

### 4️⃣ Obtener un producto por ID (GET)

```bash
curl http://localhost:8080/productos/1
```

**Respuesta:**
```json
{
  "id": 1,
  "nombre": "Laptop",
  "precio": 899.99
}
```

### 5️⃣ Actualizar un producto (PUT)

```bash
curl -X PUT http://localhost:8080/productos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Laptop Gaming",
    "precio": 1299.99
  }'
```

**Respuesta:**
```json
{
  "id": 1,
  "nombre": "Laptop Gaming",
  "precio": 1299.99
}
```

### 6️⃣ Eliminar un producto (DELETE)

```bash
curl -X DELETE http://localhost:8080/productos/1
```

**Respuesta:**
```json
{
  "mensaje": "Producto eliminado exitosamente"
}
```

## 🔍 Probar con Postman o Thunder Client

Si prefieres una interfaz gráfica, puedes usar:

- **Postman**: https://www.postman.com/
- **Thunder Client**: Extensión de VS Code
- **Insomnia**: https://insomnia.rest/

### Ejemplo de petición POST en Postman:

1. Método: `POST`
2. URL: `http://localhost:8080/productos`
3. Headers: `Content-Type: application/json`
4. Body (raw JSON):
```json
{
  "nombre": "Mouse",
  "precio": 25.99
}
```

## 📚 Conceptos Clave

### 1. **JSON Tags**
```go
type Producto struct {
    ID     int     `json:"id"`
    Nombre string  `json:"nombre" binding:"required"`
    Precio float64 `json:"precio" binding:"required,gt=0"`
}
```
- `json:"nombre"` - Nombre del campo en JSON
- `binding:"required"` - Campo obligatorio
- `binding:"gt=0"` - Mayor que 0

### 2. **Códigos HTTP**
- `200 OK` - Solicitud exitosa
- `201 Created` - Recurso creado
- `400 Bad Request` - Datos inválidos
- `404 Not Found` - Recurso no encontrado

### 3. **Estructura de Handlers**
```go
func Handler(c *gin.Context) {
    // c.Param("id") - Parámetros de URL
    // c.ShouldBindJSON() - Parsear JSON
    // c.JSON() - Enviar respuesta JSON
}
```

## 🎯 Próximos Pasos

Una vez que domines esta API, puedes:

1. ✅ Agregar **validaciones** más complejas
2. ✅ Conectar con una **base de datos** (PostgreSQL, MySQL)
3. ✅ Implementar **autenticación** con JWT
4. ✅ Agregar **paginación** a la lista de productos
5. ✅ Implementar **CORS** para frontend
6. ✅ Agregar **middleware** para logging
7. ✅ Crear **tests unitarios**

## 🐛 Errores Comunes

### Error: "package github.com/gin-gonic/gin is not in GOROOT"

**Solución:**
```bash
go mod download
```

### Error: "bind: address already in use"

**Solución:** El puerto 8080 ya está en uso. Cambia el puerto en `main.go`:
```go
router.Run(":8081")  // Usa otro puerto
```

## 📖 Documentación de Gin

- Sitio oficial: https://gin-gonic.com/
- GitHub: https://github.com/gin-gonic/gin
- Ejemplos: https://github.com/gin-gonic/examples
