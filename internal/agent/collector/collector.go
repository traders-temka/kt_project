package collector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetJSON[T any](url string, data *T) error {
	get_http, err := http.Get(url)
	defer get_http.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v \n", err)
	}
	fmt.Printf("STATUS: %v \n", get_http.Status)

	return json.NewDecoder(get_http.Body).Decode(data)
}


// func ConnectRedis() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Адрес, который мы пробросили в Docker
// 		Password: "",               // По умолчанию пароля нет
// 		DB:       0,                // Номер базы
// 	})
// }
