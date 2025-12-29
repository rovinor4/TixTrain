package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	config := &gorm.Config{}
	if os.Getenv("DEBUG_SQL") == "true" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	DB, err = gorm.Open(postgres.Open(dsn), config)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
