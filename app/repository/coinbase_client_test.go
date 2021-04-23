package repository

import (
	"log"
	"maribowman/portfolio-monitor/app/model"
	"net/http"
	"testing"
)

func TestGetHoldings(t *testing.T) {
	asset, _ := NewCoinbaseClient().GetHoldings("btc")
	log.Println(model.ToString(asset))
}

func TestAuthenticationHeaders(t *testing.T) {
	headers := generateAuthenticationHeaders(http.MethodGet, "/api?test=true")
	log.Println(headers)
}
