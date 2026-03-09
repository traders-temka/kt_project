package collector

import (
	"fmt"
	"kt_project/internal/models"
	"os"
	"time"
	"strconv"
)

type Binance struct {}

type binanceResponse struct {
	Price string `json: "price"`
}

func (b Binance) GetStat(coin string) models.Stat { //Get information from market
	url := "https://api.binance.com/api/v3/ticker/price?symbol=" + b.formatSymbol(coin)
	var resp binanceResponse
	err := GetJSON(url, &resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Dont get JSON (%v)\n", err)
	}
	price, _:= strconv.ParseFloat(resp.Price, 64)
	return models.Stat{
		Name:  coin,
		Price: price,
		Source: "Binance",
		Timedump: time.Now(),
	}
}

func (b Binance) formatSymbol(coin string) string {
    return coin + "USDT"
}
