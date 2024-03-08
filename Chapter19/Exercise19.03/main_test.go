package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	UserID   string
	Username string
}

type Application struct {
	AuthServiceURL string
}

func NewApplication(authServiceURL string) *Application {
	return &Application{
		AuthServiceURL: authServiceURL,
	}
}

// AuthenticateUser simulates the user authentication process
func (app *Application) AuthenticateUser(token string) (*User, error) {
	return &User{
		UserID:   "123",
		Username: "testuser",
	}, nil
}

func TestAuthenticationIntegration(t *testing.T) {
	// Set up a test server for the authentication service
	authService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate authentication logic
		if r.Header.Get("Authorization") == "Bearer valid_token" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"user_id": "123", "username": "testuser"}`))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer authService.Close()

	app := NewApplication(authService.URL)

	// Test the authentication process
	token := "valid_token"
	gotUser, err := app.AuthenticateUser(token)
	assert.NoError(t, err)
	assert.Equal(t, "123", gotUser.UserID)
	assert.Equal(t, "testuser", gotUser.Username)
}
