package repository

import "main/src/models"

type PostgresRepository interface {
	GetAllcoin(coins *[]models.CoinData) error
	SaveDataToDatabase(coins []struct {
		Name  string `json:"name"`
		Quote struct {
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quote"`
		LastUpdated string `json:"last_updated"`
	})
	GetAllcoinByID(coins *models.CoinData, coin_name string) error
}
