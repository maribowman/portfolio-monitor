package service

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"io/ioutil"
	"maribowman/portfolio-monitor/app/model"
	"os"
	"strconv"
)

func createBase64DonutChart(assets []model.Asset, positions []model.Holding) (string, error) {
	donut := chart.DonutChart{
		Title:        "Cryptos in Euro",
		TitleStyle:   chart.Style{},
		ColorPalette: chart.AlternateColorPalette,
		Width:        700,
		Height:       500,
		DPI:          0,
		Values:       generateDonutItems(assets, positions),
	}
	file, _ := os.Create("donut.png")
	defer file.Close()
	donut.Render(chart.PNG, file)
	content, _ := ioutil.ReadAll(bufio.NewReader(file))
	return base64.StdEncoding.EncodeToString(content), nil
}

func generateDonutItems(assets []model.Asset, positions []model.Holding) []chart.Value {
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
