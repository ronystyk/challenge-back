package services

import (
	"challenge/models"
	"strings"

	"gorm.io/gorm"
)

type StockService struct {
	db *gorm.DB
}

func NewStockService(db *gorm.DB) *StockService {
	return &StockService{db: db}
}

func (s *StockService) SaveStocks(stocks []models.Stock) error {
	batchSize := 500
	return s.db.CreateInBatches(stocks, batchSize).Error
}

func (s *StockService) GetPaginatedStocks(page, limit int, filter string) ([]models.Stock, int64, error) {
	var stocks []models.Stock
	var total int64
	var filterLower string = strings.ToLower(filter)

	result := s.db.Model(&models.Stock{}).Where("LOWER(ticker) LIKE ?", "%"+filterLower+"%").Or("LOWER(company) LIKE ?", "%"+filterLower+"%").Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	offset := (page - 1) * limit
	result = s.db.Order("time desc").Limit(limit).Offset(offset).Where("LOWER(ticker) LIKE ?", "%"+filterLower+"%").Or("LOWER(company) LIKE ?", "%"+filterLower+"%").Find(&stocks)
	if result.Error != nil {
		return nil, total, result.Error
	}

	return stocks, total, nil
}

func (s *StockService) GetRecommendations() ([]models.Stock, error) {
	var recommended []models.Stock

	result := s.db.Raw(`
        SELECT * 
		FROM stocks
		WHERE target_from > 0 AND target_to > 0
		ORDER BY (
			(
				CASE 
					WHEN rating_to = 'Strong Buy' THEN 1.4
					WHEN rating_to = 'Buy' THEN 1.3
					WHEN rating_to = 'Outperform ' THEN 1.2
					WHEN rating_to = 'Overweight' THEN 1.1
					WHEN rating_to IN ('Neutral', 'Hold') THEN 1.0
					WHEN rating_to = 'Underperform' THEN .5
					WHEN rating_to = 'Sell' THEN 0.1
					ELSE 1.0
				END
			) * ((target_to - target_from) / target_from * 100)
		) DESC
		LIMIT 12
    `).Scan(&recommended)
	if result.Error != nil {
		return nil, result.Error
	}

	return recommended, nil
}
