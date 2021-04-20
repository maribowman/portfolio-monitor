package service

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
	"log"
	"maribowman/portfolio-monitor/app/model"
	"os"
	"strconv"
)

func drawPieChart(assets []model.Asset, positions []model.Position) {
	pie := charts.NewPie()
	pie.Renderer = render.NewChartRender(pie, pie.Validate)
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Crypto Assets",
			Top:   "center",
			Left:  "center",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeChalk,
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show: true,
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show: true,
					Name: "crypto-pie",
				},
			},
		}),
	)

	pie.AddSeries("pie chart", generatePieItems(assets, positions)).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
			}),
		)
	file, err := os.Create("pie.html")
	//file, err := os.Create("pie.png")
	if err != nil {
		log.Println(err.Error())
	}


	log.Printf("data:\n%v", pie.Assets)


	if err := pie.Render(file); err != nil {
		log.Println(err.Error())
	}
}

func generatePieItems(assets []model.Asset, positions []model.Position) []opts.PieData {
	items := make([]opts.PieData, 0)
	for _, position := range positions {
		for _, asset := range assets {
			if position.Ticker == asset.Ticker {
				positionValue, _ := strconv.ParseFloat(position.Amount, 64)
				assetValue, _ := strconv.ParseFloat(asset.Amount, 64)
				items = append(items, opts.PieData{
					Name:  position.Ticker,
					Value: fmt.Sprintf("%f", positionValue*assetValue),
				})
			}
		}
	}
	return items
}
