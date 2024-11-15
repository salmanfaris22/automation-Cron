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
