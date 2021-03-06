package service

import (
	"maribowman/portfolio-monitor/app/model"
	"testing"
)

func TestCharts(t *testing.T) {
	tables := []struct {
		positions []model.Holding
		assets    []model.Asset
	}{
		{
			positions: []model.Holding{
				{Ticker: "ETH", Amount: "0.500"},
				{Ticker: "BCH", Amount: "1.000"},
				{Ticker: "BTC", Amount: "0.02011424"},
			},
			assets: []model.Asset{
				{Ticker: "ETH", Currency: "EUR", Amount: "1814.53"},
				{Ticker: "BCH", Currency: "EUR", Amount: "784.45"},
				{Ticker: "BTC", Currency: "EUR", Amount: "464766.59"},
			},
		},
	}
	for _, table := range tables {

		createBase64DonutChart(table.assets, table.positions)
		//drawPieChart(table.assets, table.positions)

	}
}
