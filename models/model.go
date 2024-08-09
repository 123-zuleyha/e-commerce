package models

import (
	"gorm.io/gorm"
)

// Product yapısı , ürünleri temsil eder.
type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"size.255,not null"`
	Description string  `json:"description" gorm:"size:500"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
}
