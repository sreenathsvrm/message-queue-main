package main

import (
	"log"
	"message-queue/config"
	"message-queue/connection"
	"message-queue/handlers"
	"message-queue/producer"
	"message-queue/repository"
	"message-queue/server"
	"message-queue/service"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("error retrieving configuration " + err.Error())
	}

	db, err := connection.ConnectDB(cfg)
	if err != nil {
		log.Fatal("databse connection error: " + err.Error())
	}

	producer := producer.ConnectBroker(cfg)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, producer)
	productHandler := handlers.NewProductService(productService)

	router := server.InitRouter(productHandler)

	// Start the server
	router.Run("localhost:8080")
}
