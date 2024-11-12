package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshua468/currency-converter/internal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("EXCHANGE_API_KEY")
	if apiKey == "" {
		log.Fatal("API key not found")
	}

	internal.StartTUI(apiKey)
}
