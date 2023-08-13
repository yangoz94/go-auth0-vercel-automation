package utils

import (
	"fmt"
	"os"
	"testing"
)

type ContainsTest struct {
	slice    []string
	city     string
	expected bool
}

func TestContains(t *testing.T) {
	tests := []ContainsTest{
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "Istanbul", true},
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "Ankara", true},
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "Izmir", true},
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "Antalya", false},
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "", false},
		{nil, "Istanbul", false},
		{nil, "", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Contains(%v, %s) should be %v", test.slice, test.city, test.expected), func(t *testing.T) {
			result := Contains(test.slice, test.city)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}


type apiKeyTest struct {
	name     string
	apiKey   string
	expected bool
}

func TestIsAPIKeyValid(t *testing.T) {
	// Set a valid API key for testing
	validAPIKey := "valid-api-key"
	os.Setenv("API_KEY", validAPIKey)

	tests := []apiKeyTest{
		{"ValidAPIKey", validAPIKey, true},
		{"InvalidAPIKey", "invalid-api-key", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAPIKeyValid(test.apiKey)
			if result != test.expected {
				t.Errorf("Expected %v for API key validity, but got %v", test.expected, result)
			}
		})
	}

	// Clean up  env var after the tests
	os.Unsetenv("API_KEY")
}
