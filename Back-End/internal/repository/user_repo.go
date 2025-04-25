package repository

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
)

type UserRepo interface {
	GetUserByEmail(email string) (*models.User, *helpers.CustomError)
	CreateUser(user *models.User) (*models.User, *helpers.CustomError)
	GetAllUsers() ([]*models.User, *helpers.CustomError)
}
