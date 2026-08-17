package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/jonestm9/economic-os/backend/internal/ingestion/bls"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading BLS API key")
	}
	apiKey := os.Getenv("BLS_API_KEY")

	httpClient := &http.Client{}

	blsClient := bls.NewClient(
		httpClient,
		apiKey,
	)

	// hardcoded response fetching unemployment data from 2026 and 2025
	response, err := blsClient.FetchData(
		[]string{"LNS14000000"},
		"2025",
		"2026",
	)

	if err != nil {
		log.Fatal(err)
	}


	for _, series := range response.Results.Series {

		fmt.Println("Series:", series.SeriesID)

		for _, observation := range series.Data {
			fmt.Printf(
				"%s %s: %s\n",
				observation.Year,
				observation.Period,
				observation.Value,
			)
		}
	}
}