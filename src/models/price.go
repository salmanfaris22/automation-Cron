package models

import (
	"time"
)

type CoinData struct {
	ID          uint      `gorm:"primary_key"`
	Name        string    `gorm:"type:varchar(100)"`
	Price       float64   `gorm:"type:decimal(20,8)"`
	LastUpdated time.Time `gorm:"type:timestamp"`
}

type Response struct {
	Data []struct {
		Name  string `json:"name"`
		Quote struct {
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quote"`
		LastUpdated string `json:"last_updated"`
	} `json:"data"`
}

type CoinMarketCapResponse struct {
	Data []CoinData `json:"data"`
}
