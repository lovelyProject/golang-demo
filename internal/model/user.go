package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string  `gorm:"unique" validate:"required,email"`
	Password  string  `gorm:"not null" validate:"required,min=8"`
	Name      string  `gorm:"not null" validate:"required"`
	Phone     *string `gorm:"unique" validate:"numeric,min=5,max=16"`
	SessionId *string `gorm:"unique"`
	Code      *string
}
