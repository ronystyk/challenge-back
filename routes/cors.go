package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupCors(router *mux.Router) (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption) {
	// Middleware de CORS
	corsAllowedOrigins := handlers.AllowedOrigins([]string{"*"})
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	return corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders
}
