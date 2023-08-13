package deployment

import (
	"auth0-vercel-script/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type Deployment struct {
	UID     string `json:"uid"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Created int64  `json:"created"`
	Source  string `json:"source,omitempty"`
	State   string `json:"state,omitempty"`
	Type    string `json:"type"`
}

type DeploymentsData struct {
	Deployments []Deployment `json:"deployments"`
}

func FetchDeploymentURLs(vercelToken string) ([]string, error) {
	if vercelToken == "" {
		return nil, errors.New("No Vercel Token provided")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.vercel.com/v6/deployments?target=production", nil)
	if err != nil {
		return nil, errors.New("Error creating request")
	}
	req.Header.Set("Authorization", "Bearer "+vercelToken)

	resp, err := client.Do(req)
	switch resp.StatusCode {
	case 400:
		return nil, errors.New("400 - Invalid query parameters")
	case 403:
		return nil, errors.New("403 - Invalid Vercel Token")
	case 404:
		return nil, errors.New("404 - Not Found")
	case 422:
		return nil, errors.New("422 - Unprocessable Entity")
	default:
		if resp.StatusCode != 200 {
			return nil, errors.New("Error fetching deployment URLs")
		}
	}

	defer resp.Body.Close()

	var result DeploymentsData
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, errors.New("Error decoding response body")
	}

	var URLs []string
	for _, deployment := range result.Deployments {
		if !utils.Contains(URLs, deployment.URL) {
			httpsURL := "https://" + deployment.URL
			URLs = append(URLs, httpsURL)
		}
	}
	return URLs, nil
}
