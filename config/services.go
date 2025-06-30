package config

import (
	"challenge/services"

	"gorm.io/gorm"
)

type Services struct {
	StockService *services.StockService
}

func InitServices(db *gorm.DB) *Services {
	return &Services{
		StockService: services.NewStockService(db),
	}
}
