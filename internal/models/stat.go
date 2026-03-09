package models

import "time"

type Stat struct {
	Name     string    `json:"symbol"`
	Price    float64   `json:"price,string"`
	Source   string    `json:"source"`
	Timedump time.Time `json:"timedump" swaggerignore:"true"`
}

type Exchange interface {
	GetStat(coin string) Stat //Get information from market
}


