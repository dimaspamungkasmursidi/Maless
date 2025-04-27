package services

import (
	"log"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/repository"
)

type TagServiceImpl struct {
	tagRepository repository.TagRepo
}

func NewTagServiceImpl(tagRepository repository.TagRepo) TagService {
	return &TagServiceImpl{
		tagRepository: tagRepository,
	}
}

func (t *TagServiceImpl) GetTagById(idTag int) (*web.TagResponse, *helpers.CustomError) {

	if idTag == 0 {
		return nil, helpers.NewCustomError("Tag ID cannot be empty", 400)
	}

	tag, err := t.tagRepository.GetTagById(idTag)
	if err != nil {
		return nil, err
	}

	log.Println("[SERVICE] Successfully retrieved tag by ID:", tag.ID)
	return &web.TagResponse{
		ID:      tag.ID,
		TagName: tag.TagName,
		UserID:  tag.UserID,
	}, nil
}

func (t *TagServiceImpl) CreateTag(tag *web.TagRequest) (*web.TagResponse, *helpers.CustomError) {

	if tag.TagName == "" {
		return nil, helpers.NewCustomError("Tag name cannot be empty", 400)
	} else if tag.UserID == 0 {
		return nil, helpers.NewCustomError("User ID cannot be empty", 400)
	}

	defineTag := &models.Tag{
		TagName: tag.TagName,
		UserID:  tag.UserID,
	}

	createTag, err := t.tagRepository.CreateTag(defineTag)
	if err != nil {
		return nil, err
	}
	log.Println("[SERVICE] Successfully created tag:", createTag.TagName)
	return &web.TagResponse{
		ID:      createTag.ID,
		TagName: createTag.TagName,
	}, nil
}

func (t *TagServiceImpl) UpdateTag(id_tag int, tag *web.TagRequest) (*web.TagResponse, *helpers.CustomError) {

	if tag.TagName == "" {
		return nil, helpers.NewCustomError("Tag name cannot be empty", 400)
	} else if id_tag == 0 {
		return nil, helpers.NewCustomError("User ID cannot be empty", 400)
	}

	if id_tag == 0 {
		return nil, helpers.NewCustomError("Tag ID cannot be empty", 400)
	}

	defineTag := &models.Tag{
		TagName: tag.TagName,
		UserID:  tag.UserID,
	}

	updateTag, err := t.tagRepository.UpdateTag(id_tag, defineTag)
	if err != nil {
		return nil, err
	}

	log.Println("[SERVICE] Successfully updated tag:", updateTag.TagName)
	return &web.TagResponse{
		ID:      updateTag.ID,
		TagName: updateTag.TagName,
		UserID:  updateTag.UserID,
	}, nil
}
