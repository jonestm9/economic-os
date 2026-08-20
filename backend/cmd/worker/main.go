package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jonestm9/economic-os/backend/internal/clients/fred"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: .env file not loaded: %v", err)
	}

	fredAPIKey := os.Getenv("FRED_API_KEY")
	if fredAPIKey == "" {
		log.Fatal("FRED_API_KEY is not set")
	}

	fredClient := fred.NewClient(fredAPIKey)

	var trackedReleases = []int{
		10,
		18,
		50,
		53,
	}
	
	for _, releaseID := range trackedReleases {
		release, err := fredClient.GetRelease(releaseID)
		if err != nil {
			log.Fatal(err)
		}
	
		fmt.Printf(
			"%d | %s | press release: %t | %s\n",
			release.ID,
			release.Name,
			release.PressRelease,
			release.Link,
		)
	}
}