package verceldeployments

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

type DeploymentURLsTest struct {
	name        string
	vercelToken string
	expectErr   bool
}

func TestFetchDeploymentURLs(t *testing.T) {
	// Load environment variables from .env.local file for testing
	err := godotenv.Load("../.././.env.local")
	if err != nil {
		t.Fatal("Error loading .env.local file")
	}

	tests := []DeploymentURLsTest{
		{
			name:        "Valid Vercel Token",
			vercelToken: os.Getenv("VERCEL_TOKEN"),
			expectErr:   false,
		},
		{
			name:        "Invalid Vercel Token",
			vercelToken: "invalid_token_hehe",
			expectErr:   true,
		},
		{
			name:        "Empty Vercel Token",
			vercelToken: "",
			expectErr:   true,
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
