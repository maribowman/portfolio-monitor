package repository

import (
	"log"
	"maribowman/portfolio-monitor/app/model"
	"testing"
)

func TestGetPrice(t *testing.T) {
	asset, _ := NewCoinbaseClient().GetPrice("btc", "usd")
	log.Println(model.ToString(asset))
}
