package handlers

import "github.com/gin-gonic/gin"

type UserHandlers interface {
	CreateUser(c *gin.Context)
}
