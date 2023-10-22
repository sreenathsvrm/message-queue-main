package models

import (
	"strings"

	"gorm.io/gorm"
)

type KafkaPayload struct {
	PublicUserID  string
	UserID        uint64
	SafeID        uint64
	ClanID        uint64
	TransactionID string
	LogTime       string
	LogDesc       string
	LogChan       string
	LogError      string
}

type QueuePayload struct {
	ProductID int
}

type ProductModel struct {
	DB *gorm.DB
}

type ImageUrls struct {
	url []string
}

func (p *ProductModel) GetProductUrls(productID int) (string, error) {
	var urls string

	err := p.DB.Raw("SELECT product_image FROM products WHERE id=?", productID).Scan(&urls).Error
	if err != nil {
		return "", err
	}

	return urls, nil
}

func (p *ProductModel) UpdateProductUrls(urls []string, productID int) error {
	productImageArray := "{" + strings.Join(urls, ",") + "}"

	err := p.DB.Exec("UPDATE products SET compressed_product_image=? WHERE id=?",productImageArray, productID).Error
	if err != nil {
		return err
	}

	return nil
}
