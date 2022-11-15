package gorm

import (
	"database/sql"

	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) repositories.PhotoRepository {
	return &photoRepository{
		db: db,
	}
}

func (p *photoRepository) GetPhotos() (*[]models.Photo, error) {
	var photos []models.Photo
	err := p.db.Find(&photos).Error
	if err != nil {
		return nil, err
	}

	if len(photos) == 0 {
		return nil, sql.ErrNoRows
	}

	return &photos, nil
}

func (p *photoRepository) UploadPhoto(photo *models.Photo) error {
	err := p.db.Create(photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *photoRepository) UpdatePhoto(photo *models.Photo) error {
	err := p.db.Save(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *photoRepository) DeletePhoto(id string) error {
	err := p.db.Delete(&models.Photo{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
