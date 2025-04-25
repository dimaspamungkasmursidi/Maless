package handlers

import (
	"net/http"
	"strconv"

	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/services"
	"github.com/gin-gonic/gin"
)

type TagHandlerImpl struct {
	TagService services.TagService
}

func NewTagHandlerImpl(tagService services.TagService) TagHandlers {
	return &TagHandlerImpl{
		TagService: tagService,
	}
}

func (t *TagHandlerImpl) GetTagByIdHandler(c *gin.Context) {

	idTag, err := strconv.Atoi(c.Param("idTag"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid tag ID"})
		return
	}

	tagResponse, customError := t.TagService.GetTagById(idTag)
	if customError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customError})
		return
	}
	tagRes := web.WebResponse{
		Code:    http.StatusOK,
		Message: "Tag retrieved successfully",
		Data:    tagResponse,
	}

	c.JSON(http.StatusOK, gin.H{"data": tagRes})
}

func (t *TagHandlerImpl) CreateTagHandler(c *gin.Context) {

	var tagRequest web.TagRequest
	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user := c.MustGet("user").(*models.User)

	tagRequest.UserID = user.ID

	tagResponse, customError := t.TagService.CreateTag(&tagRequest)
	if customError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customError})
		return
	}
	tagRes := web.WebResponse{
		Code:    http.StatusCreated,
		Message: "Tag created successfully",
		Data:    tagResponse,
	}

	c.JSON(http.StatusCreated, gin.H{"data": tagRes})
}

func (t *TagHandlerImpl) UpdateTagHandler(c *gin.Context) {

	idTag, err := strconv.Atoi(c.Param("idTag"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid tag ID"})
		return
	}

	var tagRequest web.TagRequest
	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	tagResponse, customError := t.TagService.UpdateTag(idTag, &tagRequest)
	if customError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customError})
		return
	}
	tagRes := web.WebResponse{
		Code:    http.StatusOK,
		Message: "Tag updated successfully",
		Data:    tagResponse,
	}

	c.JSON(http.StatusOK, gin.H{"data": tagRes})
}
