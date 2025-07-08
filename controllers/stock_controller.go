package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"challenge/fetcher"
	"challenge/services"
)

type StockController struct {
	Service *services.StockService
}

// Constructor
func NewStockController(service *services.StockService) *StockController {
	return &StockController{
		Service: service,
	}
}

func (c *StockController) GetStocks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))
	filter := query.Get("filter")
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	stocks, total, err := c.Service.GetPaginatedStocks(page, limit, filter)
	if err != nil {
		http.Error(w, "Error al consultar datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"page":   page,
		"limit":  limit,
		"total":  total,
		"stocks": stocks,
	})
}

func (c *StockController) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	recommended, err := c.Service.GetRecommendations()
	if err != nil {
		http.Error(w, "Error al consultar datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"recommendations": recommended,
		"message":         "✅ Recomendaciones obtenidas correctamente",
	})
}

func (c *StockController) SyncStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := fetcher.FetchAllStocks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.Service.SaveStocks(stocks)
	if err != nil {
		http.Error(w, "Error al guardar en la base de datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "✅ Datos sincronizados correctamente",
		"count":   len(stocks),
	})
}
