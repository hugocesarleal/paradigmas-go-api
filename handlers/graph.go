package handlers

import (
	"strconv"
	"strings"
	"time"

	"os"
	"trab-final/services"

	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// gera o gráfico de barras
func HandleGraph(c *gin.Context) {
	stocks := strings.Split(os.Getenv("STOCKS"), ",")

	// tempo sequencial
	startSeq := time.Now()
	services.FetchStocksSequential(stocks)
	durationSeq := time.Since(startSeq).Seconds()

	// tempo paralelo
	startPar := time.Now()
	services.FetchStocksParallel(stocks)
	durationPar := time.Since(startPar).Seconds()

	// cria o gráfico de barras
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Benchmark: Sequential vs Parallel",
			Subtitle: "Tempo em segundos para " + strconv.Itoa(len(stocks)) + " requisições",
		}),
	)

	bar.SetXAxis([]string{"Sequencial", "Paralelo"}).
		AddSeries("Tempo (s)", []opts.BarData{
			{Value: durationSeq,
				ItemStyle: &opts.ItemStyle{Color: "#1f77b4"}},
			{Value: durationPar,
				ItemStyle: &opts.ItemStyle{Color: "#ff7f0e"}},
		}).SetSeriesOptions(
		charts.WithBarChartOpts(opts.BarChart{BarWidth: "20%"}),
	)

	page := components.NewPage()
	page.AddCharts(bar)

	// renderiza diretamente na resposta HTTP
	c.Header("Content-Type", "text/html")
	page.Render(c.Writer)
}
