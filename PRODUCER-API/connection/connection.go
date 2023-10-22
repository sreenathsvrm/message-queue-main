package connection

import (
	"fmt"
	"message-queue/config"
	"message-queue/internals/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to database
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DB_HOST, cfg.DB_USER , cfg.DB_NAME, cfg.DB_PORT, cfg.DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Products{})

	return db, nil
}
