package services

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
)

type UserServices interface {
	CreateUser(userRequest *web.UserRequest) (*web.UserResponse, *helpers.CustomError)
}
