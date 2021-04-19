package service

import (
	"github.com/gin-gonic/gin"
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
	return service.coinbaseClient.GetPrice(ticker)
}

func (service *CoinbaseService) GetCrypto(context *gin.Context, coinTicker string) {
}
