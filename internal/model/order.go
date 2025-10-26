package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID   uint      `gorm:"column:user_id"`
	Products []Product `gorm:"many2many:order_products;"`
}
