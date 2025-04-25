package handlers

import "github.com/gin-gonic/gin"

type TagHandlers interface {
	GetTagByIdHandler(c *gin.Context)
	CreateTagHandler(c *gin.Context)
	UpdateTagHandler(c *gin.Context)
}
