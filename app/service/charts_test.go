package service

import (
	"maribowman/portfolio-monitor/app/model"
	"testing"
)

func TestCharts(t *testing.T) {
	tables := []struct {
		positions []model.Position
		assets    []model.Asset
	}{
		{
			positions: []model.Position{
				{Ticker: "ETH", Amount: "0.500"},
				{Ticker: "BCH", Amount: "1.000"},
				{Ticker: "BTC", Amount: "0.02011424"},
			},
			assets: []model.Asset{
				{Ticker: "ETH", Currency: "EUR", Amount: "1767.53"},
				{Ticker: "BCH", Currency: "EUR", Amount: "754.45"},
				{Ticker: "BTC", Currency: "EUR", Amount: "45866.59"},
			},
		},
	}
	for _, table := range tables {

		drawPieChart(table.assets, table.positions)

	}
}
