package main

import (
	"log"
	"net/http"

	"challenge/config"
	"challenge/routes"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	godotenv.Load()

	// Inicializar base de datos y servicios
	db := config.InitDB()
	services := config.InitServices(db)

	// Configurar rutas
	router := routes.SetupRoutes(services)

	// Configurar CORS
	corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders := routes.SetupCors(router)

	// Configurar
	log.Println("Servidor en http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders)(router))
}
