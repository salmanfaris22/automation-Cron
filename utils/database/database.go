package database

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"main/src/models"
)

var once sync.Once
var DB *gorm.DB

func GetInstancePostgres() (db *gorm.DB) {
	once.Do(func() {

		user := viper.GetString("postgres.user")
		password := viper.GetString("postgres.password")
		host := viper.GetString("postgres.host")
		port := viper.GetString("postgres.port")
		dbname := viper.GetString("postgres.dbname")
		sslmode := viper.GetString("postgres.sslmode")
		timezone := viper.GetString("postgres.timezone")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			host, user, password, dbname, port, sslmode, timezone)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Panic().Msgf("Error connecting to the database at %s:%s/%s", host, port, dbname)
			log.Info().Msgf("Connection details: %s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
		}

		if err := db.AutoMigrate(&models.CoinData{}); err != nil {
			log.Panic().Msgf("Error during auto-migration: %v", err)
		}

		DB = db
	})
	return DB
}
