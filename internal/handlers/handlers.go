package handlers

import (
	auth0callbacks "auth0-vercel-script/internal/auth0-callbacks"
	vercelDeployments "auth0-vercel-script/internal/vercel-deployments"
	"auth0-vercel-script/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/auth0/go-auth0/management"
	"github.com/joho/godotenv"
)

func UpdateCallbackURLsHandler(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}

	// Check if the request has a valid API key
	apiKey := r.Header.Get("API-Key")
	if apiKey == "" {
		http.Error(w, "Missing API key", http.StatusUnauthorized)
		return
	} else if !utils.IsAPIKeyValid(apiKey) {
		http.Error(w, "Invalid API key", http.StatusUnauthorized)
		return
	}

	// Parse the request body
	var requestData struct {
		Auth0Domain       string `json:"auth0_domain"`
		Auth0ClientID     string `json:"auth0_client_id"`
		Auth0ClientSecret string `json:"auth0_client_secret"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate Auth0 credentials
	if requestData.Auth0Domain == "" || requestData.Auth0ClientID == "" || requestData.Auth0ClientSecret == "" {
		http.Error(w, "Missing Auth0 credentials", http.StatusBadRequest)
		return
	}

	domain := requestData.Auth0Domain
	clientID := requestData.Auth0ClientID
	clientSecret := requestData.Auth0ClientSecret

	auth0API, err := management.New(
		domain,
		management.WithClientCredentials(context.Background(), clientID, clientSecret),
	)
	if err != nil {
		http.Error(w, "Invalid credentials: "+err.Error(), http.StatusUnauthorized) // Return Unauthorized status
		return
	}

	urls, err := vercelDeployments.FetchDeploymentURLs()
	if err != nil {
		http.Error(w, "Failed to fetch deployment URLs: "+err.Error(), http.StatusBadGateway)
		return
	}

	currentURLs, err := auth0callbacks.GetCurrentCallbackURLs(auth0API, clientID)
	if err != nil {
		http.Error(w, "Failed to fetch current callback URLs: One or more Auth0 credentials are invalid/incorrect ", http.StatusBadGateway)
		return
	}

	message, newURLs, err := auth0callbacks.UpdateClientCallbacks(auth0API, clientID, urls)
	if err != nil {
		http.Error(w, "Failed to update callback URLs: Something went wrong", http.StatusBadGateway)
		return
	}

	// Combine the newly added URLs with the current URLs
	updatedURLs := append(currentURLs, newURLs...)

	// Prepare the response payload
	response := struct {
		Message        string   `json:"message"`
		NewlyAddedURLs []string `json:"newly_added_urls,omitempty"`
		UpdatedURLs    []string `json:"current_urls,omitempty"`
	}{
		Message:        message,
		NewlyAddedURLs: newURLs,
		UpdatedURLs:    updatedURLs,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)

	// Marshal the response as JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
