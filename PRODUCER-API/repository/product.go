package repository

import (
	"message-queue/internals/models"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(product models.Products) (int, error)
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (p *productRepository) AddProduct(product models.Products) (int, error) {
	productID := 0
	productImageArray := "{" + strings.Join(product.ProductImage, ",") + "}"
	compressedProductImageArray := "{" + strings.Join(product.CompressedProductImage, ",") + "}"

	err := p.DB.Raw("INSERT INTO products(product_name, product_description, product_image, product_price, compressed_product_image) VALUES(?, ?, ?, ?, ?) RETURNING id", product.ProductName, product.ProductDescription, productImageArray, product.ProductPrice, compressedProductImageArray).Scan(&productID).Error
	if err != nil {
		return 0, err
	}

	return productID, nil
}
