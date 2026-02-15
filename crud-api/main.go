package main

import (
	"crud-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear el router de Gin
	router := gin.Default()

	// Configurar las rutas
	routes.SetupRoutes(router)

	// Mensaje de inicio
	log.Println("Servidor iniciado en http://localhost:8080")

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
