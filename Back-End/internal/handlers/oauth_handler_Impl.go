package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/oauth"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OauthHandlersImpl struct {
	UserService services.UserServices
}

func NewOauthHandlers(userService services.UserServices) *OauthHandlersImpl {
	return &OauthHandlersImpl{
		UserService: userService,
	}
}

func (h *OauthHandlersImpl) HandleGoogleLogin(c *gin.Context) {
	url := oauth.GoogleOauthConfig.AuthCodeURL(oauth.OauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *OauthHandlersImpl) HandleGoogleCallback(c *gin.Context) {
	state := c.DefaultQuery("state", "")
	if state != oauth.OauthStateString {
		log.Println("Invalid oauth state")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid oauth state"})
		return
	}

	code := c.DefaultQuery("code", "")
	token, err := oauth.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Code exchange failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Code exchange failed"})
		return
	}

	idToken := token.Extra("id_token").(string)
	c.SetCookie("user-session", idToken, 3600, "/", "localhost", false, true)

	// Pindahkan logika getUserInfo ke method terpisah
	userInfo, customErr := h.GetUserInfo(token)
	if customErr != nil {
		log.Println("Failed getting user info:", customErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErr})
		return
	}

	googleID := userInfo["id"].(string)
	email := userInfo["email"].(string)
	name := userInfo["name"].(string)
	picture := userInfo["picture"].(string)

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

	c.Redirect(http.StatusFound, "http://localhost:3000/dashboard")
}

func (h *OauthHandlersImpl) GetUserInfo(token *oauth2.Token) (map[string]interface{}, *helpers.CustomError) {
	client := oauth.GoogleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, helpers.NewCustomError("Failed to get user info", http.StatusInternalServerError)
	}
	defer response.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, helpers.NewCustomError("Failed to parse user info", http.StatusInternalServerError)
	}

	return userInfo, nil
}
