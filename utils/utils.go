package utils

import "os"

func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func IsAPIKeyValid(apiKey string) bool {
	validAPIKey := os.Getenv("API_KEY")
	return apiKey == validAPIKey
}
