package service

import (
	"github.com/gin-gonic/gin"
	"maribowman/portfolio-monitor/app/model"
)

type CoinbaseService struct {
	coinbaseClient  model.FinanceClient
	signalMessenger model.MessengerClient
}

type Wiring struct {
	FinanceClient model.FinanceClient
	Messenger     model.MessengerClient
}

func NewCoinbaseService(wiring *Wiring) model.FinanceService {
	return &CoinbaseService{
		coinbaseClient:  wiring.FinanceClient,
		signalMessenger: wiring.Messenger,
	}
}

func (service *CoinbaseService) ProcessAsset(ticker, currency string) (model.Asset, error) {
	//holdings := []model.Holding{
	//	{Ticker: "ETH", Amount: "0.500"},
	//	{Ticker: "BCH", Amount: "1.000"},
	//	{Ticker: "BTC", Amount: "0.02011424"},
	//}
	return service.coinbaseClient.GetPrice(ticker, currency)

	//donutChart, _ := createBase64PieChart(assets, holdings)
	//donutChart, _ := createBase64DonutChart(assets, positions)

	//log.Println(donutChart)

	//return model.Asset{}, nil
}

func (service *CoinbaseService) GetCrypto(context *gin.Context, coinTicker string) {
}
