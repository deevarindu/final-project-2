package models

import "time"

type User struct {
	Id        *int      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"not null;uniqueIndex" form:"username" valid:"required~username is required"`
	Email     string    `json:"email" gorm:"not null;uniqueIndex" form:"email" valid:"required~email is required, email~invalid email format"`
	Password  string    `json:"password" gorm:"not null" form:"password" valid:"required~password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age       int       `json:"age" valid:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
