package helpers

import (
	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
)

func ModelUserCreate(user *models.User) *web.UserResponse {
	return &web.UserResponse{
		ID:        user.ID,
		GoogleId:  user.GoogleId,
		Name:      user.Name,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
	}
}
