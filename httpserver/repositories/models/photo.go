package models

import "time"

type Photo struct {
	Id        *int      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null" form:"title" valid:"required~title is required"`
	Caption   string    `json:"caption" gorm:"not null" form:"caption" valid:"required~caption is required"`
	PhotoUrl  string    `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required~photo_url is required"`
	UserId    int       `json:"user_id" gorm:"not null" form:"user_id" valid:"required~user_id is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
