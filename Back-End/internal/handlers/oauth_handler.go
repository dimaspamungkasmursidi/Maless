package handlers

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OauthHandlers interface {
	HandleGoogleLogin(c *gin.Context)
	HandleGoogleCallback(c *gin.Context)
	GetUserInfo(token *oauth2.Token) (map[string]interface{}, *helpers.CustomError)
}
