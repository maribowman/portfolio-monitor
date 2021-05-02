package model

import (
	"encoding/json"
)

type CoinbaseWrapper struct {
	Asset Asset `json:"data"`
}

type Asset struct {
	Ticker   string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Holding struct {
	Ticker string `json:"ticker"`
	ISIN   string `json:"isin"`
	Amount string `json:"amount"`
}

type Transaction struct {
	ISIN        string  `json:"isin,omitempty"`
	Ticker      string  `json:"ticker" binding:"required"`
	ActualPrice float32 `json:"actualPrice" binding:"required"`
	TotalPrice  float32 `json:"totalPrice" binding:"required"`
	Currency    string  `json:"currency,omitempty"`
	Date        string  `json:"date,omitempty"`
	Comment     string  `json:"comment,omitempty"`
}

func ToString(model interface{}) string {
	jsonBytes, err := json.Marshal(model)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

type FinanceService interface {
	ProcessAsset(ticker, currency string) (Asset, error)
}

type FinanceClient interface {
	GetPrice(ticker, currency string) (Asset, error)
	GetHoldings(ticker string) (string, error)
}
