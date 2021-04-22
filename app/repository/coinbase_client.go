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

// GetPrice Coinbase API: https://developers.coinbase.com/api/v2?shell#prices
func (client *CoinbaseClient) GetPrice(ticker, currency string) (model.Asset, error) {
	var response model.CoinbaseWrapper
	details := RequestDetails{
		Protocol: "https",
		BaseUrl:  "api.coinbase.com",
		Path:     fmt.Sprintf("/v2/prices/%s-%s/spot", ticker, currency),
	}
	if err := client.restClient.getData(details, &response); err != nil {
		return model.Asset{}, err
	}
	return response.Asset, nil
}

func (client *CoinbaseClient) GetPortfolio(ticker string) (string, error) {
	var response string
	details := RequestDetails{
		Protocol: "https",
		BaseUrl:  "api.coinbase.com",
		Path:     "/v2/accounts",
		Query:    nil,
		Headers:  authenticationHeaders(http.MethodGet, "/v2/accounts"),
		Body:     "",
	}
	if err := client.restClient.getData(details, &response); err != nil {
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
