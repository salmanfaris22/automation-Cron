package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main/src/service"
)

func GetCoinData(ctx *gin.Context) {
	coinData, err := service.GetAllCoins()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch coin data",
		})
		return
	}

	ctx.JSON(http.StatusOK, coinData)
}