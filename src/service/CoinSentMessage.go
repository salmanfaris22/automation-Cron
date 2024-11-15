package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"main/src/models"
	"main/src/repository"
)

func SentMessage() {
	marketcapURL := viper.GetString("coinmarketcap.url")
	APIKey := viper.GetString("coinmarketcap.api")
	APIName := viper.GetString("coinmarketcap.name")

	req, err := http.NewRequest("GET", marketcapURL, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	req.Header.Add(APIName, APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	var response models.Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	repo := repository.NewPostgresRepo()
	repo.SaveDataToDatabase(response.Data)
}
