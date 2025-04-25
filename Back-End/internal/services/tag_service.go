package services

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
)

type TagService interface {
	GetTagById(idTag int) (*web.TagResponse, *helpers.CustomError)
	CreateTag(tag *web.TagRequest) (*web.TagResponse, *helpers.CustomError)
	UpdateTag(id_tag int, tag *web.TagRequest) (*web.TagResponse, *helpers.CustomError)
}
