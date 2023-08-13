
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

func UpdateClientCallbacks(auth0API *management.Management, clientID string, urls []string) error {
	client, err := auth0API.Client.Read(context.Background(), clientID)
	if err != nil {
		return err
	}

	missingURLs := FindMissingURLs(client, urls)
	if len(missingURLs) > 0 {
		updatedCallbacks := append(*client.Callbacks, missingURLs...)
		err := UpdateCallbacks(auth0API, clientID, updatedCallbacks)
		if err != nil {
			return err
		}
		PrintMissingURLs(missingURLs)
	} else {
		fmt.Println("No missing URLs found.")
	}

	return nil
}
