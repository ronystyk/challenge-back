package main

import (
	"log"
	"net/http"

	"challenge/config"
	"challenge/routes"
	"challenge/services"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Inicializar base de datos y servicios
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	srvs := services.InitServices(db)

	// Configurar rutas
	router := routes.SetupRoutes(srvs)

	// Configurar CORS
	corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders := routes.SetupCors(router)

	// Configurar
	log.Println("Servidor en http://localhost:8080")
	err = http.ListenAndServe(":8080", handlers.CORS(corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders)(router))
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
