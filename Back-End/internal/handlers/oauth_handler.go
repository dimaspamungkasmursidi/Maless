package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/oauth"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func HandleGoogleLogin(c *gin.Context) {
	url := oauth.GoogleOauthConfig.AuthCodeURL(oauth.OauthStateString)
	fmt.Println("URL:", url)
	c.Redirect(http.StatusTemporaryRedirect, url)

}

func HandleGoogleCallback(c *gin.Context) {
	fmt.Println("Callback hit")
	state := c.Query("state")
	if state != oauth.OauthStateString {
		log.Println("Invalid oauth state, expected:", oauth.OauthStateString, "got:", state)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid oauth state"})
		return
	}

	code := c.Query("code")
	log.Println("Authorization code received:", code)

	token, err := oauth.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Code exchange failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Code exchange failed"})
		fmt.Println("Token received:", token)
		return
	}

	userInfo, err := GetUserInfo(token)
	if err != nil {
		log.Println("Failed getting user info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed getting user info"})
		return
	}

	fmt.Println("Access token received:", token.AccessToken)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"user":    userInfo,
	})
}

func GetUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := oauth.GoogleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: status code %d", response.StatusCode)
	}

	var userInfo map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}
	fmt.Println("User Info:", userInfo)
	return userInfo, nil

}
