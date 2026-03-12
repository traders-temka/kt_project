package collector

import (
	"fmt"
	"backend/internal/models"
	"os"
	"time"
	"strconv"
)

type Bybit struct {}


func (b Bybit) GetStat(baseCoin string, quoteCoin string) (models.Stat, error) { //Get information from market
	url := fmt.Sprintf("https://api.bybit.com/v5/market/tickers?category=spot&symbol=%s%s", baseCoin, quoteCoin)
	var resp struct {
    	Result struct {
    		List []struct {
    	        BidPrice string `json:"bid1Price"` //JSON structure
    	        AskPrice string `json:"ask1Price"`
    	    } `json:"list"`
    	} `json:"result"`
	}

	err := GetJSON(url, &resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Dont get JSON (%v)\n", err)
		return models.Stat{}, fmt.Errorf("No data")
	}
	if len(resp.Result.List) == 0 { return models.Stat{}, fmt.Errorf("No data") }
	bidprice, _:= strconv.ParseFloat(resp.Result.List[0].BidPrice, 64)
	askprice, _:= strconv.ParseFloat(resp.Result.List[0].AskPrice, 64)
	return models.Stat{
		Base: baseCoin,
		Quote: quoteCoin,
		AskPrice: askprice,
		BidPrice: bidprice,
		Source: "Bybit",
		Timedump: time.Now(),
	}, nil
}
