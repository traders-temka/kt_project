package models

import "time"

type Stat struct {
	Base     string    `json:"base"`
	Quote    string    `json:"quote"`
	AskPrice float64   `json:"ask_price"`
	BidPrice float64   `json:"bid_price"`
	Source   string    `json:"source"`
	Timedump time.Time `json:"timedump" swaggerignore:"true"`
}

type Exchange interface {
	GetStat(baseCoin string, quoteCoin string) (Stat, error) //Get information from market
}
