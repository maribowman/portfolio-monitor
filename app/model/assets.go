package model

type CoinbaseWrapper struct {
	Asset Asset `json:"data"`
}

type Asset struct {
	Ticker   string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Position struct {
	Ticker string `json:"ticker"`
	Isin   string `json:"isin"`
	Amount string `json:"amount"`
}

type FinanceService interface {
	ProcessAsset(ticker string) (Asset, error)
}

type FinanceClient interface {
	GetPrice(ticker string) (Asset, error)
}

type Messenger interface {
	Push(price float32, recipient string) error
}
