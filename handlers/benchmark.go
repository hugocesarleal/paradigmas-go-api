package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"trab-final/services"

	"github.com/gin-gonic/gin"
)

// define um tipo para a resposta de cada ação
type StockResult struct {
	Stock string      `json:"stock"`
	Data  interface{} `json:"data"`
}

// define um tipo para benchmark
type BenchmarkResult struct {
	Sequential float64 `json:"sequential"`
	Parallel   float64 `json:"parallel"`
}

func HandleSequential(c *gin.Context) {
	stocks := strings.Split(os.Getenv("STOCKS"), ",")
	start := time.Now()

	var results []StockResult
	for _, stock := range stocks {
		raw := services.FetchStock(stock)
		// se o services.FetchStock retorna JSON cru, você pode desserializar:
		var doc interface{}
		if err := json.Unmarshal([]byte(raw), &doc); err != nil {
			// em caso de erro de parsing, guarda a string bruta
			doc = raw
		}
		results = append(results, StockResult{
			Stock: stock,
			Data:  doc,
		})
	}

	elapsed := time.Since(start).Seconds()

	c.IndentedJSON(http.StatusOK, gin.H{
		"mode":     "sequential",
		"duration": elapsed,
		"results":  results,
	})
}

func HandleParallel(c *gin.Context) {
	stocks := strings.Split(os.Getenv("STOCKS"), ",")
	start := time.Now()

	raws := services.FetchStocksParallel(stocks)

	var results []StockResult
	for i, raw := range raws {
		var doc interface{}
		if err := json.Unmarshal([]byte(raw), &doc); err != nil {
			doc = raw
		}
		results = append(results, StockResult{
			Stock: stocks[i],
			Data:  doc,
		})
	}

	elapsed := time.Since(start).Seconds()

	c.IndentedJSON(http.StatusOK, gin.H{
		"mode":     "parallel",
		"duration": elapsed,
		"results":  results,
	})
}

func HandleBenchmark(c *gin.Context) {
	stocks := strings.Split(os.Getenv("STOCKS"), ",")

	// sequencial
	t0 := time.Now()
	services.FetchStocksSequential(stocks)
	durationSeq := time.Since(t0).Seconds()

	// paralelo
	t1 := time.Now()
	services.FetchStocksParallel(stocks)
	durationPar := time.Since(t1).Seconds()

	bench := BenchmarkResult{
		Sequential: durationSeq,
		Parallel:   durationPar,
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"benchmark": bench,
	})
}
