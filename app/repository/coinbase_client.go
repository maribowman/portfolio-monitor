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
	var response model.CoinbaseWrapper
	if err := client.restClient.getData("api.coinbase.com", "/v2", fmt.Sprintf("/prices/%s/spot", ticker), nil, nil, &response); err != nil {
		return model.Asset{}, err
	}
	return response.Asset, nil
}
