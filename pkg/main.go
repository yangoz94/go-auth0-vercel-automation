package main

import (
	"auth0-vercel-script/internal/deployment"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	start := time.Now()

	// Load environment variables from .env.local file
	err := godotenv.Load(".env.local"); if err != nil {
		log.Fatal("Error loading .env.local file")
	}
	
	vercelToken := os.Getenv("VERCEL_TOKEN")
	urls, err := deployment.FetchDeploymentURLs(vercelToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("URLs: ", urls)
	elapsed := time.Since(start).Seconds()
	fmt.Printf("Done! - Script took %.2fs to run\n", elapsed)
}

