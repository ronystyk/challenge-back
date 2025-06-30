package routes

import (
	"challenge/config"
	"challenge/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(services *config.Services) *mux.Router {
	// Inicializar controladores
	stockController := controllers.NewStockController(services.StockService)

	router := mux.NewRouter()
	router.HandleFunc("/stocks", stockController.GetStocks).Methods("GET")
	router.HandleFunc("/recommendations", stockController.GetRecommendations).Methods("GET")
	router.HandleFunc("/sync-stocks", stockController.SyncStocks).Methods("POST")

	return router
}
