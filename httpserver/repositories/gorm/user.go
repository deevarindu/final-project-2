package gorm

import (
	"database/sql"

	"github.com/deevarindu/final-project-2/httpserver/repositories"
	"github.com/deevarindu/final-project-2/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUsers() (*[]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, sql.ErrNoRows
	}

	return &users, nil
}

func (u *userRepository) GetUser(id string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) CreateUser(user *models.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateUser(user *models.User) error {
	err := u.db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(id string) error {
	err := u.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
