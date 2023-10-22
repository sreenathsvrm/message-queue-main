package connection

import (
	"consumer/config"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connection to database
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DB_HOST, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PORT, cfg.DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// connect to kafka
func ConnectBroker(cfg *config.Config) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.BROKER_ADDR},
		Topic:   cfg.BROKER_TOPIC,
		GroupID: "product-group",
	})

	log.Println("broker init")

	return reader
}
