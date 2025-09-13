package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Type        string
	Price       float64
	Images      pq.StringArray `gorm:"type:text[];default:'{}'"`
}
