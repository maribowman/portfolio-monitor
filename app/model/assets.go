package model

type Asset struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
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
