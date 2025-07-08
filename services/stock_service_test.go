package services

import (
	"testing"

	"challenge/config"
	"challenge/models"
)

func TestSaveStocks(t *testing.T) {
	db, err := config.InitTestDB(t)
	if err != nil {
		t.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// Instanciar el servicio de stock
	stockService := NewStockService(db)

	// creamos el mock de los stocks
	stocks := []models.Stock{
		{
			Ticker:     "TEST",
			Company:    "Test Inc",
			Brokerage:  "FakeBroker",
			Action:     "initiated by",
			RatingFrom: "Sell",
			RatingTo:   "Buy",
			TargetFrom: 10.0,
			TargetTo:   20.0,
		},
	}

	// Llamamos al método SaveStocks del servicio
	err = stockService.SaveStocks(stocks)
	if err != nil {
		t.Errorf("Error al guardar los stocks: %v", err)
	}

	// Verificar que los stocks se hayan guardado correctamente
	var savedStocks []models.Stock
	if err := db.Find(&savedStocks).Error; err != nil {
		t.Errorf("Error al recuperar los stocks guardados: %v", err)
	}

	if len(savedStocks) != len(stocks) {
		t.Errorf("Se esperaban %d stocks guardados, pero se encontraron %d", len(stocks), len(savedStocks))
	}
}

func TestGetPaginatedStocks(t *testing.T) {
	db, err := config.InitTestDB(t)
	if err != nil {
		t.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// Instanciar el servicio de stock
	stockService := NewStockService(db)

	// Cargar datos de prueba
	stocks := []models.Stock{
		{
			Ticker:     "TEST1",
			Company:    "Test Inc 1",
			Brokerage:  "FakeBroker",
			Action:     "initiated by",
			RatingFrom: "Sell",
			RatingTo:   "Buy",
			TargetFrom: 10.0,
			TargetTo:   20.0,
		},
		{
			Ticker:     "TEST2",
			Company:    "Test Inc 2",
			Brokerage:  "FakeBroker",
			Action:     "initiated by",
			RatingFrom: "Hold",
			RatingTo:   "Sell",
			TargetFrom: 15.0,
			TargetTo:   25.0,
		},
	}
	if err := db.Create(&stocks).Error; err != nil {
		t.Fatalf("Error al cargar los datos de prueba: %v", err)
	}

	page, limit := 1, 10
	filter := ""

	resultStocks, total, err := stockService.GetPaginatedStocks(page, limit, filter)
	if err != nil {
		t.Errorf("Error al obtener los stocks paginados: %v", err)
	}

	if len(resultStocks) != len(stocks) {
		t.Errorf("Se esperaban %d stocks, pero se encontraron %d", len(stocks), len(resultStocks))
	}

	if total != int64(len(stocks)) {
		t.Errorf("Se esperaba un total de %d stocks, pero se obtuvo %d", len(stocks), total)
	}
}

func TestGetRecommendations(t *testing.T) {
	db, err := config.InitTestDB(t)
	if err != nil {
		t.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// Instanciar el servicio de stock
	stockService := NewStockService(db)

	// Cargar datos de prueba
	stocks := []models.Stock{
		{
			Ticker:     "TEST1",
			Company:    "Test Inc 1",
			Brokerage:  "FakeBroker",
			Action:     "initiated by",
			RatingFrom: "Sell",
			RatingTo:   "Buy",
			TargetFrom: 10.0,
			TargetTo:   20.0,
		},
	}
	if err := db.Create(&stocks).Error; err != nil {
		t.Fatalf("Error al cargar los datos de prueba: %v", err)
	}

	recommendations, err := stockService.GetRecommendations()
	if err != nil {
		t.Errorf("Error al obtener las recomendaciones: %v", err)
	}

	if len(recommendations) == 0 {
		t.Error("Se esperaban recomendaciones, pero no se encontraron")
	}
}
