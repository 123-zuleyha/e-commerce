package models

import (
	"gorm.io/gorm"
	 "github.com/123-zuleyha/e-commerce/config"

)

 

// Product yapısı , ürünleri temsil eder.
type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"size.255,not null"`
	Description string  `json:"description" gorm:"size:500"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
}

// GetAllProducts tüm ürünleri döndürür
func GetAllProducts() ([]Product, error) {
	var products []Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductByID belirli bir ID'ye sahip ürünü döndürür
func GetProductByID(id string) (*Product, error) {
	var product Product

	if err := config.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// CreateProduct yeni bir ürün oluştur
func CreateProduct(product *Product) error {
	return config.DB.Create(product).Error
}


func UpdateProduct(id string, product *Product) error {
	return config.DB.Model(&Product{}).Where("id = ?", id).Updates(product).Error
}

// DeleteProduct belirli bir ID'ye sahip ürünü siler
func DeleteProduct(id string) error {
	return config.DB.Delete(&Product{}, id).Error
}