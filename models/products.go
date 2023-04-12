package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// ProductCategory create model and use the GORM model to add functions for CRUD
type ProductCategory struct {
	*gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"unique;size:128;not null"`
	Description string `gorm:"size:1024"`
}

func (c *ProductCategory) String() string {
	return c.Name
}

// GetAllProductCategories Get  all categories.
func GetAllProductCategories() ([]ProductCategory, error) {
	var categories []ProductCategory
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// Product - model for storing information about products
type Product struct {
	*gorm.Model
	ID          uint    `gorm:"primary_key"`
	Name        string  `gorm:"size:128;not null"`
	Description string  `gorm:"size:2048"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Quantity    uint    `gorm:"default:0"`
	Image       string  `gorm:"size:256"`
	CategoryID  uint    `gorm:"not null"`
	Category    *ProductCategory
}

// The String method returns a string representation of the product for output to the console.
func (p Product) String() string {
	return fmt.Sprintf("Product: %s || Category: %s", p.Name, p.Category.Name)
}

func GetAllProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductsByCategory(categoryID uint) ([]Product, error) {
	var products []Product
	err := db.Where("category_id = ?", categoryID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
