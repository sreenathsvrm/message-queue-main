package service

import (
	"encoding/json"
	"errors"
	"message-queue/internals/models"
	"message-queue/producer"
	"message-queue/repository"
)

type ProductService interface {
	AddProduct(product models.Products) error
}

type productService struct {
	productRepo repository.ProductRepository
	producer    producer.Producer
}

func NewProductService(productRepo repository.ProductRepository, producer producer.Producer) ProductService {
	return &productService{
		productRepo: productRepo,
		producer:    producer,
	}
}

func (p *productService) AddProduct(product models.Products) error {
	// add product in database
	productID, err := p.productRepo.AddProduct(product)
	if err != nil {
		return err
	}

	if productID == 0 {
		return errors.New("product could not be created ")
	}

	productWrapper := producer.QueuePayload{
		ProductID : productID,
	}

	// serialize payload
	productPayload, err := json.Marshal(productWrapper)
	if err != nil {
		return err
	}

	// send product id to message broker
	err = p.producer.Produce(productPayload)
	if err != nil {
		return err
	}

	return nil
}
