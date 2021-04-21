package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"maribowman/portfolio-monitor/app/model"
)

type CoinbaseService struct {
	coinbaseClient  model.FinanceClient
	signalMessenger model.Messenger
}

type Wiring struct {
	FinanceClient model.FinanceClient
	Messenger     model.Messenger
}

func NewCoinbaseService(wiring *Wiring) model.FinanceService {
	return &CoinbaseService{
		coinbaseClient:  wiring.FinanceClient,
		signalMessenger: wiring.Messenger,
	}
}

func (service *CoinbaseService) ProcessAsset(ticker string) (model.Asset, error) {
	holdings := []model.Holding{
		{Ticker: "ETH", Amount: "0.500"},
		{Ticker: "BCH", Amount: "1.000"},
		{Ticker: "BTC", Amount: "0.02011424"},
	}
	var assets []model.Asset
	for _, holding := range holdings {
		asset, _ := service.coinbaseClient.GetPrice(holding.Ticker, "EUR")
		assets = append(assets, asset)
	}

	donutChart, _ := createBase64PieChart(assets, holdings)
	//donutChart, _ := createBase64DonutChart(assets, positions)

	log.Println(donutChart)

	return model.Asset{}, nil
}

func (service *CoinbaseService) GetCrypto(context *gin.Context, coinTicker string) {
}
