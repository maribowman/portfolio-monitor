package repository

import (
	"fmt"
	"maribowman/portfolio-monitor/app/model"
)

type CoinbaseClient struct {
	restClient *RestClient
}

func NewCoinbaseClient() model.FinanceClient {
	return &CoinbaseClient{
		restClient: NewRestClient(),
	}
}

func (client *CoinbaseClient) GetPrice(ticker string) (model.Asset, error) {
	var response model.Asset
	if err := client.restClient.getData("api.coinbase.com", "/v2", fmt.Sprintf("/prices/%s/spot", ticker), nil, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}
