package collector

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetJSON[T any](url string, data *T) error {
	get_http, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("network error: %w", err)
	}

	defer get_http.Body.Close()

	if get_http.StatusCode != http.StatusOK {
		return fmt.Errorf("api returned bad status: %s", get_http.Status)
	}

	if err := json.NewDecoder(get_http.Body).Decode(data); err != nil {
		return fmt.Errorf("failed to decode json: %w", err)
	}

	return nil
}


// func ConnectRedis() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Адрес, который мы пробросили в Docker
// 		Password: "",               // По умолчанию пароля нет
// 		DB:       0,                // Номер базы
// 	})
// }
