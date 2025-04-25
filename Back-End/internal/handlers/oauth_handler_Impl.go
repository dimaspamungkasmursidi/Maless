package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/oauth"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

var (
	store = sessions.NewCookieStore([]byte("your-secret-key"))
)

type OauthHandlersImpl struct {
	UserService services.UserServices
}

func NewOauthHandlers(userService services.UserServices) OauthHandlers {
	return &OauthHandlersImpl{
		UserService: userService,
	}
}

func (h *OauthHandlersImpl) HandleGoogleLogin(c *gin.Context) {
	url := oauth.GoogleOauthConfig.AuthCodeURL(oauth.OauthStateString)
	fmt.Println("URL:", url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *OauthHandlersImpl) HandleGoogleCallback(c *gin.Context) {
	state := c.DefaultQuery("state", "")
	if state != oauth.OauthStateString {
		log.Println("Invalid oauth state, expected:", oauth.OauthStateString, "got:", state)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid oauth state"})
		return
	}

	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not found"})
		return
	}
	log.Println("Authorization code received:", code)

	token, err := oauth.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Code exchange failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Code exchange failed"})
		return
	}
	log.Println("Access Token received:", token.AccessToken)

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.Println("ID token not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID token missing"})
		return
	}
	log.Println("ID Token:", idToken)

	c.SetCookie("user-session", idToken, 3600, "/", "localhost", false, true)

	userInfo, customErr := h.GetUserInfo(token)
	if customErr != nil {
		log.Println("Failed getting user info:", customErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErr})
		return
	}
	log.Println("User Info:", userInfo)

	googleID, ok := userInfo["id"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Google ID"})
		return
	}

	email, ok := userInfo["email"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid email"})
		return
	}

	name, ok := userInfo["name"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid name"})
		return
	}

	picture, ok := userInfo["picture"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid picture URL"})
		return
	}

	user := &web.UserRequest{
		GoogleId:  googleID,
		Email:     email,
		Name:      name,
		AvatarUrl: picture,
	}

	_, customErr = h.UserService.CreateUser(user)
	if customErr != nil {
		log.Println("Failed to save user:", customErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErr})
		return
	}

	c.Redirect(http.StatusFound, "http://localhost:3000")
}

func (h *OauthHandlersImpl) GetUserInfo(token *oauth2.Token) (map[string]interface{}, *helpers.CustomError) {
	client := oauth.GoogleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, helpers.NewCustomError("Failed to get user info from Google", http.StatusInternalServerError)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, helpers.NewCustomError("Google API returned non-200 status", response.StatusCode)
	}

	var userInfo map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, helpers.NewCustomError("Failed to parse user info", http.StatusInternalServerError)
	}

	if _, ok := userInfo["id"]; !ok {
		return nil, helpers.NewCustomError("Google ID not found in response", http.StatusInternalServerError)
	}
	if _, ok := userInfo["email"]; !ok {
		return nil, helpers.NewCustomError("Email not found in response", http.StatusInternalServerError)
	}

	log.Printf("User Info Received: %+v", userInfo)
	return userInfo, nil
}
