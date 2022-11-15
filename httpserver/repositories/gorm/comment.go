package gorm

import (
	"database/sql"

	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repositories.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) GetComments() (*[]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Find(&comments).Error
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, sql.ErrNoRows
	}

	return &comments, nil
}

func (c *commentRepository) CreateComment(comment *models.Comment) error {
	err := c.db.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *commentRepository) UpdateComment(comment *models.Comment) error {
	err := c.db.Save(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *commentRepository) DeleteComment(id string) error {
	err := c.db.Delete(&models.Comment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
