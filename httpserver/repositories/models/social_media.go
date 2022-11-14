package models

type SocialMedia struct {
	Id             *int   `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"not null" form:"name" valid:"required~name is required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" form:"social_media_url" valid:"required~social_media_url is required"`
	UserId         int    `json:"user_id" gorm:"not null" form:"user_id" valid:"required~user_id is required"`
}
