package models

import "time"

type Comment struct {
	Id        *int      `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"user_id" gorm:"not null" form:"user_id" valid:"required~user_id is required"`
	PhotoId   int       `json:"photo_id" gorm:"not null" form:"photo_id" valid:"required~photo_id is required"`
	Message   string    `json:"message" gorm:"not null" form:"message" valid:"required~message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
