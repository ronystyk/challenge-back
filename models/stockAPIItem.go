package models

import (
	"time"
)

type StockAPIItem struct {
	Ticker     string    `json:"ticker"`
	Company    string    `json:"company"`
	Brokerage  string    `json:"brokerage"`
	Action     string    `json:"action"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	TargetFrom string    `json:"target_from"` // ← viene como "$1,234.56"
	TargetTo   string    `json:"target_to"`   // ← viene como "$1,234.56"
	Time       time.Time `json:"time"`
}
