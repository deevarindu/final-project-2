package gorm

import (
	"database/sql"

	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) repositories.SocialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (s *socialMediaRepository) GetSocialMedias() (*[]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := s.db.Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}

	if len(socialMedias) == 0 {
		return nil, sql.ErrNoRows
	}

	return &socialMedias, nil
}

func (s *socialMediaRepository) AddSocialMedia(socialMedia *models.SocialMedia) error {
	err := s.db.Create(socialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *socialMediaRepository) UpdateSocialMedia(socialMedia *models.SocialMedia) error {
	err := s.db.Save(&socialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *socialMediaRepository) DeleteSocialMedia(id string) error {
	err := s.db.Delete(&models.SocialMedia{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
