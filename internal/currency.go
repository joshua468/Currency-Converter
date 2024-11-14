package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RateResponse is the struct for API response.
type RateResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

// FetchAllRates fetches the exchange rates using the API key.
func FetchAllRates(apiKey string) (map[string]float64, error) {
	url := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/USD?apikey=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rateData RateResponse
	if err := json.NewDecoder(resp.Body).Decode(&rateData); err != nil {
		return nil, err
	}

	return rateData.Rates, nil
}

// ConvertCurrency performs the currency conversion calculation.
func ConvertCurrency(amount, fromRate, toRate float64) float64 {
	return (amount / fromRate) * toRate
}
