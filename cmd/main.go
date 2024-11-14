package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/joshua468/currency-converter/internal"
	"github.com/spf13/cobra"
)

var ConvertCommand = &cobra.Command{
	Use:   "convert",
	Short: "Convert between two currencies",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch the API key from the environment variable
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			fmt.Println("API_KEY is missing. Please set it in your .env file.")
			return
		}

		// Fetch rates from the API using the API key
		rates, err := internal.FetchAllRates(apiKey)
		if err != nil {
			fmt.Println("Error fetching exchange rates:", err)
			return
		}

		var amountStr string
		fmt.Print("Enter amount: ")
		fmt.Scanln(&amountStr)

		// Convert the input amount to a float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		// Get the list of currencies
		currencies := getCurrencyList(rates)
		for i, currency := range currencies {
			fmt.Printf("%d. %s\n", i+1, currency)
		}

		// Let the user select from and to currencies
		var fromCurrencyIndex, toCurrencyIndex int
		fmt.Print("Select From Currency (number): ")
		fmt.Scanln(&fromCurrencyIndex)
		fromCurrency := currencies[fromCurrencyIndex-1]

		fmt.Print("Select To Currency (number): ")
		fmt.Scanln(&toCurrencyIndex)
		toCurrency := currencies[toCurrencyIndex-1]

		// Get the exchange rates for the selected currencies
		fromRate := rates[fromCurrency]
		toRate := rates[toCurrency]

		// Perform the conversion
		converted := internal.ConvertCurrency(amount, fromRate, toRate)

		// Show the result
		result := fmt.Sprintf("%.2f %s = %.2f %s", amount, fromCurrency, converted, toCurrency)
		fmt.Println(result)
	},
}

func init() {
	// Load the .env file
	err := godotenv.Load("../internal/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	ConvertCommand.Flags().String("api-key", "", "API key for fetching exchange rates")
}

func main() {
	if err := ConvertCommand.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}

// getCurrencyList returns a list of currency codes from the provided rates map.
func getCurrencyList(rates map[string]float64) []string {
	currencies := make([]string, 0, len(rates))
	for currency := range rates {
		currencies = append(currencies, currency)
	}
	return currencies
}
