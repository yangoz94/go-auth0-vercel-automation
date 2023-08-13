package deployment

import (
	"os"
	"testing"
	"github.com/joho/godotenv"
)

func TestFetchDeploymentURLs(t *testing.T) {
	// Load environment variables from .env.local file for testing
	godotenv.Load("../../.env.local")

	tests := []struct {
		name       string
		vercelToken string
		expectErr  bool
	}{
		{
			name:       "Valid Vercel Token",
			vercelToken: os.Getenv("VERCEL_TOKEN"),
			expectErr:  false,
		},
		{
			name:       "Invalid Vercel Token",
			vercelToken: "invalid_token_hehe",
			expectErr:  true,
		},
		{
			name:       "Empty Vercel Token",
			vercelToken: "",
			expectErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			URLs, err := FetchDeploymentURLs(test.vercelToken)
			if test.expectErr {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Error fetching deployment URLs: %v", err)
				}
				if len(URLs) == 0 {
					t.Error("No deployment URLs fetched")
				}
			}
		})
	}
}