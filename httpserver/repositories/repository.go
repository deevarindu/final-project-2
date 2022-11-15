package repositories

import "github.com/deevarindu/final-project-2/httpserver/repositories/models"

type UserRepository interface {
	GetUsers() (*[]models.User, error)
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type PhotoRepository interface {
	GetPhotos() (*[]models.Photo, error)
	UploadPhoto(photo *models.Photo) error
	UpdatePhoto(photo *models.Photo) error
	DeletePhoto(id string) error
}

type CommentRepository interface {
	GetComments() (*[]models.Comment, error)
	CreateComment(comment *models.Comment) error
	UpdateComment(comment *models.Comment) error
	DeleteComment(id string) error
}

type SocialMediaRepository interface {
	GetSocialMedias() (*[]models.SocialMedia, error)
	AddSocialMedia(socialMedia *models.SocialMedia) error
	UpdateSocialMedia(socialMedia *models.SocialMedia) error
	DeleteSocialMedia(id string) error
}
