package handlers

import (
	"net/http"
	"strconv"

	"crud-api/models"

	"github.com/gin-gonic/gin"
)

// ListarProductos - GET /productos
// Retorna todos los productos
func ListarProductos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"productos": models.Productos,
		"total":     len(models.Productos),
	})
}

// ObtenerProducto - GET /productos/:id
// Retorna un producto específico por ID
func ObtenerProducto(c *gin.Context) {
	// Obtener el ID de los parámetros de la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Buscar el producto
	for _, producto := range models.Productos {
		if producto.ID == id {
			c.JSON(http.StatusOK, producto)
			return
		}
	}

	// Si no se encuentra, retornar 404
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Producto no encontrado",
	})
}

// CrearProducto - POST /productos
// Crea un nuevo producto
func CrearProducto(c *gin.Context) {
	var nuevoProducto models.Producto

	// Bind JSON al struct y validar
	if err := c.ShouldBindJSON(&nuevoProducto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Asignar ID automático
	nuevoProducto.ID = models.SiguienteID
	models.SiguienteID++

	// Agregar a la lista
	models.Productos = append(models.Productos, nuevoProducto)

	// Retornar el producto creado con código 201
	c.JSON(http.StatusCreated, nuevoProducto)
}

// ActualizarProducto - PUT /productos/:id
// Actualiza un producto existente
func ActualizarProducto(c *gin.Context) {
	// Obtener el ID de los parámetros
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Bind del JSON
	var productoActualizado models.Producto
	if err := c.ShouldBindJSON(&productoActualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Buscar y actualizar el producto
	for i, producto := range models.Productos {
		if producto.ID == id {
			// Mantener el ID original
			productoActualizado.ID = id
			models.Productos[i] = productoActualizado

			c.JSON(http.StatusOK, productoActualizado)
			return
		}
	}

	// Si no se encuentra
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Producto no encontrado",
	})
}

// EliminarProducto - DELETE /productos/:id
// Elimina un producto
func EliminarProducto(c *gin.Context) {
	// Obtener el ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Buscar y eliminar el producto
	for i, producto := range models.Productos {
		if producto.ID == id {
			// Eliminar del slice
			models.Productos = append(models.Productos[:i], models.Productos[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"mensaje": "Producto eliminado exitosamente",
			})
			return
		}
	}

	// Si no se encuentra
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Producto no encontrado",
	})
}
