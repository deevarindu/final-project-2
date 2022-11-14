package config

import (
	"fmt"

	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

func CreateConnectionGORM() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{})

	return db, nil
}
