package handlers

import (
	"net/http"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandlersImpl struct {
	UserService services.UserServices
}

func NewUserHandlersImpl(userService services.UserServices) UserHandlers {
	return &UserHandlersImpl{UserService: userService}
}

func (u *UserHandlersImpl) CreateUser(c *gin.Context) {

	var userRequest web.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userResponse, customError := u.UserService.CreateUser(&userRequest)
	if customError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customError})
		return
	}

	WebResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Message: "User created successfully",
		Data:    userResponse,
	}

	c.JSON(http.StatusCreated, gin.H{"data": WebResponse})
}
