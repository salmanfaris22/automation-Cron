package service

import (
	"main/src/models"
	"main/src/repository"
)

func GetAllCoins() ([]models.CoinData, error) {
	repo := repository.NewPostgresRepo()
	var coins []models.CoinData
	err := repo.GetAllcoin(&coins)
	if err != nil {
		return nil, err
	}
	return coins, nil

}

func GetAllCoinsById(coin_name string) (models.CoinData, error) {
	repo := repository.NewPostgresRepo()
	var coins models.CoinData
	err := repo.GetAllcoinByID(&coins, coin_name)
	if err != nil {
		return coins, err
	}
	return coins, nil

}
