// pkg/main.go
package main

import (
	auth0callbacks "auth0-vercel-script/internal/auth0-callbacks"
	"auth0-vercel-script/internal/deployment"
	"context"
	"log"
	"os"

	"github.com/auth0/go-auth0/management"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env.local file for testing
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}

	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
	clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")

	auth0API, err := management.New(
		domain,
		management.WithClientCredentials(context.Background(), clientID, clientSecret),
	)
	if err != nil {
		log.Fatalf("Failed to initialize the Auth0 management API client: %+v", err)
	}

	urls, err := deployment.FetchDeploymentURLs()
	if err != nil {
		log.Fatal(err)
	}

	err = auth0callbacks.UpdateClientCallbacks(auth0API, clientID, urls)
	if err != nil {
		log.Fatalf("Failed to update client callbacks: %+v", err)
	}
}

