package services

import (
	"gorm.io/gorm"
)

type Services struct {
	StockService *StockService
}

func InitServices(db *gorm.DB) *Services {
	return &Services{
		StockService: NewStockService(db),
	}
}
