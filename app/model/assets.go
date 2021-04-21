package model

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

type FinanceService interface {
	ProcessAsset(ticker string) (Asset, error)
}

type FinanceClient interface {
	GetPrice(ticker, currency string) (Asset, error)
}

type Messenger interface {
	Push(holding Holding, recipient string) error
}
