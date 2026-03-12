package collector

import (
	"fmt"
	"backend/internal/models"
	"os"
	"time"
	"strconv"
)

type Binance struct {}

func (b Binance) GetStat(baseCoin string, quoteCoin string) (models.Stat, error) { //Get information from market
	url := "https://api.binance.com/api/v3/ticker/bookTicker?symbol=" + baseCoin + quoteCoin

	var resp struct {
		Bidprice string `json:"bidPrice"`
		Askprice string `json:"askPrice"`
	}

	err := GetJSON(url, &resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Dont get JSON (%v)\n", err)
		return models.Stat{}, fmt.Errorf("No data")
	}

	bidprice, _:= strconv.ParseFloat(resp.Bidprice, 64)
	askprice, _:= strconv.ParseFloat(resp.Askprice, 64)
	return models.Stat{
		Base: baseCoin,
		Quote: quoteCoin,
		AskPrice: askprice,
		BidPrice: bidprice,
		Source: "Binance",
		Timedump: time.Now(),
	}, nil
}

