package models

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ProductName            string   `gorm:"not null" json:"product_name"`
	ProductDescription     string   `gorm:"not null" json:"product_description"`
	ProductImage           []string `gorm:"type:text[]" json:"product_image"`
	ProductPrice           float64  `gorm:"not null" json:"product_price"`
	CompressedProductImage []string `gorm:"type:text[]" json:"compressed_product_image"`
}
