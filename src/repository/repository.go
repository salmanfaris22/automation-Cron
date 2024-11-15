package repository

import (
	"log"
	"time"

	"main/src/models"
	"main/utils/database"
)

type PostgresRepo struct{}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{}
}

func (r *PostgresRepo) GetAllcoin(coins *[]models.CoinData) error {
	db := database.GetInstancePostgres()
	return db.Find(coins).Error
}

func (r *PostgresRepo) SaveDataToDatabase(coins []struct {
	Name  string `json:"name"`
	Quote struct {
		USD struct {
			Price float64 `json:"price"`
		} `json:"USD"`
	} `json:"quote"`
	LastUpdated string `json:"last_updated"`
}) {
	db := database.GetInstancePostgres()
	for _, coin := range coins {
		lastUpdated, err := time.Parse(time.RFC3339, coin.LastUpdated) // -> YYYY-MM-DDTHH:MM:SSZ (e.g., 2024-11-15T13:45:30Z)
		if err != nil {
			log.Println("Error parsing last updated time:", err)
			continue
		}
		coinData := models.CoinData{
			Name:        coin.Name,
			Price:       coin.Quote.USD.Price,
			LastUpdated: lastUpdated,
		}
		var existingCoin models.CoinData
		if err := db.Where("name = ?", coin.Name).First(&existingCoin).Error; err == nil {
			existingCoin.Price = coin.Quote.USD.Price
			existingCoin.LastUpdated = lastUpdated
			if err := db.Save(&existingCoin).Error; err != nil {
				log.Println("Error updating data:", err)
			}
		} else {
			if err := db.Create(&coinData).Error; err != nil {
				log.Println("Error inserting data:", err)
			}
		}
	}
	log.Printf("coin chart updated")
}
