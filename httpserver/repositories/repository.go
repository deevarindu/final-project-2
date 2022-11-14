package repositories

import "github.com/deevarindu/final-project-2/httpserver/repositories/models"

type UserRepository interface {
	GetUsers() (*[]models.User, error)
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}
