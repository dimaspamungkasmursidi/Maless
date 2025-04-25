package oauth

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
	OauthStateString  string
)

func init() {

	clientId := helpers.GetEnv("CLIENTID")
	clientSecret := helpers.GetEnv("CLIENTSECRET")
	stateString := helpers.GetEnv("STATESTRING")

	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8081/oauth/callback",
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	OauthStateString = stateString
}
