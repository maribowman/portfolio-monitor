package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"maribowman/portfolio-monitor/app/model"
	"net/http"
	"strconv"
	"time"
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

func (client *CoinbaseClient) GetPortfolio(ticker string) (string, error) {
	var response string
	headers := authenticationHeaders(http.MethodGet, "/v2/accounts")
	if err := client.restClient.getData("api.coinbase.com", "/v2/accounts", "", headers, nil, &response); err != nil {
		return "", err
	}
	return response, nil
}

func authenticationHeaders(method, apiPath string) map[string]interface{} {
	timestamp := time.Now().UTC().Unix()
	apiKey := ""
	apiSecret := ""

	// id := ""

	sigHash := hmac.New(sha256.New, []byte(apiSecret))
	sigHash.Write([]byte(strconv.FormatInt(timestamp, 10) + method + apiPath))

	return map[string]interface{}{
		"CB-ACCESS-KEY":       apiKey,
		"CB-ACCESS-SIGN":      hex.EncodeToString(sigHash.Sum(nil)),
		"CB-ACCESS-TIMESTAMP": timestamp,
	}
}
