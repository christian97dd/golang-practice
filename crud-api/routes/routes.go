package routes

import (
	"crud-api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(router *gin.Engine) {
	// Ruta de bienvenida
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "¡Bienvenido al CRUD API de Productos!",
			"versión": "1.0",
			"endpoints": gin.H{
				"GET /productos":       "Listar todos los productos",
				"GET /productos/:id":   "Obtener un producto por ID",
				"POST /productos":      "Crear un nuevo producto",
				"PUT /productos/:id":   "Actualizar un producto",
				"DELETE /productos/:id": "Eliminar un producto",
			},
		})
	})

	// Grupo de rutas para productos
	productosRoutes := router.Group("/productos")
	{
		productosRoutes.GET("", handlers.ListarProductos)       // Listar todos
		productosRoutes.GET("/:id", handlers.ObtenerProducto)   // Obtener uno
		productosRoutes.POST("", handlers.CrearProducto)        // Crear
		productosRoutes.PUT("/:id", handlers.ActualizarProducto) // Actualizar
		productosRoutes.DELETE("/:id", handlers.EliminarProducto) // Eliminar
	}
}
