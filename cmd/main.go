package main

import (
	"kt_project/cmd/server"
	"kt_project/internal/agent"
	"kt_project/internal/models"
	"log"
	"time"
)

// @title           Crypto Metrics API
// @version         1.0
// @description     Сервис для сбора и хранения котировок.
// @host            localhost:8080
// @BasePath        /
func main() {
	//
	// 	repo := repository.NewRedisStorage("localhost:6379", "", 0)
	// 	h := &handlers.Handler{Repo: repo}

	go server.Run()

	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		var stat models.Stat

		url := "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT"

		agent.GetPrice(url, &stat)

		// agent sends data on local server by HTTP
		err := agent.SendStat("http://localhost:8080/update", &stat)
		if err != nil {
			log.Printf("Send error: %v", err)
		}
	}
}
