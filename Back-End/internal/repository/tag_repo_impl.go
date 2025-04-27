package repository

import (
	"context"
	"errors"
	"log"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TagRepoImpl struct {
	db *pgxpool.Pool
}

func NewTagRepoImpl(db *pgxpool.Pool) TagRepo {
	return &TagRepoImpl{db: db}
}
func (t *TagRepoImpl) GetTagById(idTag int) (*models.Tag, *helpers.CustomError) {

	ctx := context.Background()
	tag := &models.Tag{}
	query := `SELECT id, tag_name, user_id FROM "tag" WHERE id = $1`
	err := t.db.QueryRow(ctx, query, idTag).Scan(&tag.ID, &tag.TagName, &tag.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("[REPOSITORY] Failed to get tag by ID:", idTag)
			return nil, helpers.NewCustomError("Tag not found", 404)
		}
		log.Println("[REPOSITORY] Error getting tag by ID:", err)
		return nil, helpers.NewCustomError("Failed to get tag: "+err.Error(), 500)
	}
	log.Println("[REPOSITORY] Successfully retrieved tag by ID:", tag.ID)
	return tag, nil
}

func (t *TagRepoImpl) CreateTag(tag *models.Tag) (*models.Tag, *helpers.CustomError) {

	ctx := context.Background()
	newTag := &models.Tag{}
	query := `INSERT INTO "tag" (tag_name, user_id) VALUES ($1, $2) RETURNING id, tag_name, user_id`

	err := t.db.QueryRow(ctx, query, tag.TagName, tag.UserID).Scan(&newTag.ID, &newTag.TagName, &newTag.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("[REPOSITORY] Failed to Create tag:")
			return nil, helpers.NewCustomError("Failed to create tag: "+err.Error(), 500)
		}
	}
	log.Println("[REPOSITORY] Successfully created tag:", newTag.TagName)
	return tag, nil
}

func (t *TagRepoImpl) UpdateTag(id_tag int, tag *models.Tag) (*models.Tag, *helpers.CustomError) {

	ctx := context.Background()
	query := `UPDATE "tag" SET tag_name = $1 WHERE id = $2 RETURNING id, tag_name, user_id`
	err := t.db.QueryRow(ctx, query, tag.TagName, id_tag).Scan(&tag.ID, &tag.TagName, &tag.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("[REPOSITORY] Failed to update tag:", err)
			return nil, helpers.NewCustomError("Tag not found", 404)
		}
		log.Println("[REPOSITORY] Error updating tag:", err)
		return nil, helpers.NewCustomError("Failed to update tag: "+err.Error(), 500)
	}
	log.Println("[REPOSITORY] Successfully updated tag:", tag.TagName)
	return tag, nil
}
