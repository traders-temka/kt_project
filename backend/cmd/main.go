package main

import (
	"backend/cmd/agent"
	"backend/cmd/server"
	"backend/internal/agent/collector"
	"backend/internal/models"
	"time"
)

// @title           Crypto Monitoring API
// @version         1.0
// @description     This is a sample crypto metrics server.
// @host            localhost:8080
// @BasePath        /
func main() {
	//
	// 	repo := repository.NewRedisStorage("localhost:6379", "", 0)
	// 	h := &handlers.Handler{Repo: repo}
	myExchanges := []models.Exchange{
		collector.Binance{},
		collector.Bybit{},
		// Kraken
		//CoinBase
	}
	targetCoins := []string{"BTC", "ETH", "DOGE"} //можно с командной строки

	go server.Run()

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		agent.RunAgent(myExchanges, targetCoins)
	}
}
