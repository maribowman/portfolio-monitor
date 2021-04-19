package service

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"maribowman/portfolio-monitor/app/model"
	"os"
	"strconv"
)

func drawPieChart(assets []model.Asset, positions []model.Position) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Crypto Assets"}),
		//pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Crypto Assets"}),
		//	charts.WithToolboxOpts(opts.Toolbox{
		//		Show: false,
		//		Feature: &opts.ToolBoxFeature{
		//			SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
		//				Show:  false,
		//				Type:  "png",
		//				Name:  "crypto-pie",
		//				Title: "Crypto Assets",
		//			},
		//		},
		//	}),
	)

	pie.AddSeries("pie chart", generatePieItems(assets, positions)).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
			}),
		)
	file, _ := os.Create("pie.html")
	//file, _ := os.Create("pie.png")
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
