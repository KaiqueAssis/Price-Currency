package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExchangeRateResponse struct {
	Result          string             `json:"result"`
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func main() {

	apiKey := ""
	baseCurrency := "AUD"

	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, baseCurrency)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer requisição", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}

	var data ExchangeRateResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	fmt.Println("cotaçãodo "+baseCurrency+" em BRL:", data.ConversionRates["BRL"])

}
