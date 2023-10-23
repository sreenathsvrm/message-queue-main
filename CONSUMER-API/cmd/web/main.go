package main

import (
	"consumer/config"
	"consumer/connection"
	"consumer/models"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// consumer dependency
type application struct {
	products models.ProductModel
	reader   *kafka.Reader
}

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("error reading config " + err.Error())
	}

	// connect to database
	DB, err := connection.ConnectDB(cfg)
	if err != nil {
		log.Fatal("error reading config " + err.Error())
	}

	// connect to broker
	reader := connection.ConnectBroker(cfg)

	// add dependencies
	app := &application{
		products: models.ProductModel{DB: DB},
		reader:   reader,
	}

	// call consumer
	app.consume(context.Background())
}
