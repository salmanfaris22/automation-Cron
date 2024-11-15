package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"main/config"
	"main/route"
	"main/src/controllers"
)

func main() {
	config.InitConfig()
	router := gin.Default()
	controllers.RunCron()
	route.SetupRoutes(router)
}
