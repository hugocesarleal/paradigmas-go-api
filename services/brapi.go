package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func FetchStock(symbol string) string {
	apiKey := os.Getenv("BRAPI_API_KEY")
	url := fmt.Sprintf("%s%s", os.Getenv("BRAPI_URL"), symbol)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("Erro ao criar requisição: %v", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprintf("Erro na requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Erro: Status Code %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func FetchStocksParallel(symbols []string) []string {
	var wg sync.WaitGroup
	results := make([]string, len(symbols))

	for i, symbol := range symbols {
		wg.Add(1)
		go func(i int, sym string) {
			defer wg.Done()
			results[i] = FetchStock(sym)
		}(i, symbol)
	}

	wg.Wait()
	return results
}

func FetchStocksSequential(symbols []string) []string {
	results := []string{}
	for _, symbol := range symbols {
		result := FetchStock(symbol)
		results = append(results, result)
	}
	return results
}
