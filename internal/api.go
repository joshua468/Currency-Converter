package internal

import (
	"encoding/json"
	"net/http"
)

type RateResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

func FetchAllRates(apiKey string) (map[string]float64, error) {
	url := "https://api.exchangerate-api.com/v4/latest/USD" // Base currency set to USD
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
