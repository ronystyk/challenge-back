package controllers

import (
	"challenge/config"
	"challenge/services"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"challenge/models"

	"github.com/joho/godotenv"
)

func TestGetStocks(t *testing.T) {
	db, err := config.InitTestDB(t)
	if err != nil {
		t.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// Cargar datos de prueba
	db.Create(&models.Stock{
		Ticker:     "TEST",
		Company:    "Test Inc",
		Brokerage:  "FakeBroker",
		Action:     "initiated by",
		RatingFrom: "Sell",
		RatingTo:   "Buy",
		TargetFrom: 10.0,
		TargetTo:   20.0,
	})

	// Inyección de dependencias del servicio
	stockService := services.NewStockService(db)

	// Inyección en el controlador
	handler := NewStockController(stockService)

	req := httptest.NewRequest(http.MethodGet, "/stocks?page=1&limit=10", nil)
	w := httptest.NewRecorder()

	handler.GetStocks(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
		// Detalles del error
		body, _ := io.ReadAll(resp.Body)
		t.Errorf("response body: %s", body)
	}
}

func TestGetRecommendations(t *testing.T) {
	db, err := config.InitTestDB(t)
	if err != nil {
		t.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// Cargar datos de prueba
	db.Create(&models.Stock{
		Ticker:     "TEST",
		Company:    "Test Inc",
		Brokerage:  "FakeBroker",
		Action:     "initiated by",
		RatingFrom: "Sell",
		RatingTo:   "Buy",
		TargetFrom: 10.0,
		TargetTo:   20.0,
	})

	// Inyección de dependencias del servicio
	stockService := services.NewStockService(db)

	// Inyección en el controlador
	handler := NewStockController(stockService)

	req := httptest.NewRequest(http.MethodGet, "/recommendations", nil)
	w := httptest.NewRecorder()

	handler.GetRecommendations(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
		// Detalles del error
		body, _ := io.ReadAll(resp.Body)
		t.Errorf("response body: %s", body)
	}
}

func TestSyncStocks(t *testing.T) {
	err := godotenv.Load("../.env") // Ajusta la ruta si es necesario
	if err != nil {
		t.Error("⚠️  No se pudo cargar el archivo .env (puede ser normal en algunos entornos)")
	}
	db, err := config.InitTestDB(t)
	if err != nil {
		t.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// Servicio de stock
	stockService := services.NewStockService(db)

	// Inyección en el controlador
	handler := NewStockController(stockService)

	req := httptest.NewRequest(http.MethodPost, "/sync-stocks", nil)
	w := httptest.NewRecorder()

	handler.SyncStocks(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
		// Detalles del error
		body, _ := io.ReadAll(resp.Body)
		t.Errorf("response body: %s", body)
	}
}
