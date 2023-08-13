package auth0callbacks

import (
	"context"
	"fmt"

	"github.com/auth0/go-auth0/management"
)

func UpdateCallbacks(auth0API *management.Management, clientID string, callbacks []string) error {
	updates := &management.Client{
		Callbacks: &callbacks,
	}
	return auth0API.Client.Update(context.Background(), clientID, updates)
}

//TO-DO: Add tests for this function
func GetCurrentCallbackURLs(auth0API *management.Management, clientID string) ([]string, error) {
	client, err := auth0API.Client.Read(context.Background(), clientID)
	if err != nil {
		return nil, err
	}
	return *client.Callbacks, nil
}

func IsURLPresent(url string, callbacks []string) bool {
	for _, callbackURL := range callbacks {
		if url == callbackURL {
			return true
		}
	}
	return false
}

func FindMissingURLs(client *management.Client, urls []string) []string {
	missingURLs := make([]string, 0)
	for _, url := range urls {
		if !IsURLPresent(url, *client.Callbacks) {
			missingURLs = append(missingURLs, url)
		}
	}
	return missingURLs
}

func PrintMissingURLs(missingURLs []string) {
	if len(missingURLs) > 0 {
		fmt.Println("Added missing URLs to Auth0 allowed callbacks list:")
		for i, url := range missingURLs {
			fmt.Printf("%d - %s\n", i+1, url)
		}
	}
}

//TO-DO: Add tests for this function
func UpdateClientCallbacks(auth0API *management.Management, clientID string, urls []string) (string, []string, error) {
	client, err := auth0API.Client.Read(context.Background(), clientID)
	if err != nil {
		return "", nil, err
	}

	missingURLs := FindMissingURLs(client, urls)
	if len(missingURLs) > 0 {
		updatedCallbacks := append(*client.Callbacks, missingURLs...)
		err := UpdateCallbacks(auth0API, clientID, updatedCallbacks)
		if err != nil {
			return "", nil, err
		}
		PrintMissingURLs(missingURLs)
	} else {
		return "No missing URLs found for the client - Your callback URLs are already in sync!", nil, nil
	}

	return "Successfully updated callback URLs for the client", missingURLs, nil
}
