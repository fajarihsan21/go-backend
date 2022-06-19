package models

import (
	"time"
)

type User struct {
	IdUser    uint      `gorm:"primaryKey" json:"id_user"`
	Username  string    `json:"username"`
	Email     string    `json:"email" validate:"email"`
	Password  string    `json:"password" validate:"required,min=6"`
	Name      string    `json:"name" validate:"required"`
	Image     string    `json:"image"`
	Birthdate string    `json:"birthdate" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User
