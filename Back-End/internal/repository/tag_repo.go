package repository

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
)

type TagRepo interface {
	GetTagById(idTag int) (*models.Tag, *helpers.CustomError)
	CreateTag(tag *models.Tag) (*models.Tag, *helpers.CustomError)
	UpdateTag(id_tag int, tag *models.Tag) (*models.Tag, *helpers.CustomError)
}
