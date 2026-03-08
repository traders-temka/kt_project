package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kt_project/internal/models"
	"net/http"
	"os"
)

func GetStat(url string, data *models.Stat) { //Get information from market
	get_http, err := http.Get(url)
	defer get_http.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v \n", err)
	}
	fmt.Printf("STATUS: %v \n", get_http.Status)

	err = json.NewDecoder(get_http.Body).Decode(data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: JSON was not decode | %v\n", err)
	}
}

func SendStat(url string, data *models.Stat) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data stat: %w", err)
	}

	// making post request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// checking obj status for 201 (created)
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error server response: %s", resp.Status)
	}

	fmt.Printf("[AGENT] successed sending data. status: %s\n", resp.Status)
	return nil
}

// func ConnectRedis() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Адрес, который мы пробросили в Docker
// 		Password: "",               // По умолчанию пароля нет
// 		DB:       0,                // Номер базы
// 	})
// }
