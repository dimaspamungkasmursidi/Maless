package oauth

import (
	"errors"
	"log"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
	OauthStateString  string
	isInitialized     bool
)

func InitOAuthConfig(redirectURL string) error {
	clientId := helpers.GetEnv("CLIENTID")
	if clientId == "" {
		return errors.New("CLIENTID environment variable is empty")
	}

	clientSecret := helpers.GetEnv("CLIENTSECRET")
	if clientSecret == "" {
		return errors.New("CLIENTSECRET environment variable is empty")
	}

	stateString := helpers.GetEnv("STATESTRING")
	if stateString == "" {
		stateString = "default_random_state" // fallback value
	}

	log.Println("[AUTH] Initializing OAuth configuration...")
	log.Println("[AUTH] Client ID:", clientId)
	log.Println("[AUTH] Client Secret:", "[REDACTED]")
	log.Println("[AUTH] State String:", stateString)

	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  redirectURL,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	OauthStateString = stateString
	isInitialized = true

	return nil
}

// GetConfig mengembalikan konfigurasi OAuth yang sudah diinisialisasi
func GetConfig() (*oauth2.Config, error) {
	if !isInitialized {
		return nil, errors.New("OAuth config not initialized")
	}
	return GoogleOauthConfig, nil
}

// GetStateString mengembalikan state string
func GetStateString() (string, error) {
	if !isInitialized {
		return "", errors.New("OAuth config not initialized")
	}
	return OauthStateString, nil
}
