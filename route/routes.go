package route

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"main/src/controllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/v1")
	{
		api.GET("/coinmarketcap", controllers.GetCoinData)
		api.GET("/coinmarketcapById", controllers.GetCoinDataById)
	}
	router.Run(viper.GetString("server.port"))
}
