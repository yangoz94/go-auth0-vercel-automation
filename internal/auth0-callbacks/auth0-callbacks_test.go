package auth0callbacks

import (
	"context"
	"testing"

	"github.com/auth0/go-auth0/management"
	"github.com/stretchr/testify/assert"
)

type MockManagementAPI struct {
	ClientUpdateErr   error
	ClientReadErr     error
	CallbacksToUpdate []string
}

func (m *MockManagementAPI) UpdateClient(ctx context.Context, clientID string, client *management.Client) error {
	if m.ClientUpdateErr != nil {
		return m.ClientUpdateErr
	}
	m.CallbacksToUpdate = *client.Callbacks
	return nil
}

func (m *MockManagementAPI) ReadClient(ctx context.Context, clientID string) (*management.Client, error) {
	if m.ClientReadErr != nil {
		return nil, m.ClientReadErr
	}
	return &management.Client{
		Callbacks: &m.CallbacksToUpdate,
	}, nil
}

func TestIsURLPresent(t *testing.T) {
	callbacks := []string{"https://oguzhanyangoz.com/callback1", "https://oguzhanyangoz.com/callback2"}

	assert.True(t, IsURLPresent("https://oguzhanyangoz.com/callback1", callbacks))
	assert.False(t, IsURLPresent("https://oguzhanyangoz.com/callback3", callbacks))
}

func TestFindMissingURLs(t *testing.T) {
	client := &management.Client{
		Callbacks: &[]string{"https://oguzhanyangoz.com/callback1", "https://oguzhanyangoz.com/callback2"},
	}

	urls := []string{"https://oguzhanyangoz.com/callback1", "https://oguzhanyangoz.com/callback3", "https://oguzhanyangoz.com/callback4"}

	missingURLs := FindMissingURLs(client, urls)

	assert.Len(t, missingURLs, 2)
	assert.Contains(t, missingURLs, "https://oguzhanyangoz.com/callback3")
	assert.Contains(t, missingURLs, "https://oguzhanyangoz.com/callback4")
}

func TestPrintMissingURLs(t *testing.T) {
	missingURLs := []string{"https://oguzhanyangoz.com/callback3", "https://oguzhanyangoz.com/callback4"}

	// The following tests whether it will rturn any erorr (it shouldn't)
	PrintMissingURLs(missingURLs)
}
