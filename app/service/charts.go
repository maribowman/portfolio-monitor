package service

import (
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"maribowman/portfolio-monitor/app/model"
	"os"
	"strconv"
)

func drawDonutChart(assets []model.Asset, positions []model.Position) {
	donut := chart.DonutChart{
		Title:        "Crypto Assets",
		TitleStyle:   chart.Style{},
		ColorPalette: chart.AlternateColorPalette,
		Width:        512,
		Height:       512,
		DPI:          0,
		Values:       generateDonutItems(assets, positions),
	}

	file, _ := os.Create("donut.png")
	defer file.Close()
	donut.Render(chart.PNG, file)
}

func generateDonutItems(assets []model.Asset, positions []model.Position) []chart.Value {
	items := make([]chart.Value, 0)
	for _, position := range positions {
		for _, asset := range assets {
			if position.Ticker == asset.Ticker {
				positionValue, _ := strconv.ParseFloat(position.Amount, 64)
				assetValue, _ := strconv.ParseFloat(asset.Amount, 64)
				value := positionValue * assetValue
				items = append(items, chart.Value{
					Label: fmt.Sprintf("%s: %.2f", position.Ticker, value),
					Value: value,
				})
			}
		}
	}
	return items
}
