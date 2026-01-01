package pkg

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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

// Paginate : make function pagination for gorm
// usage: scopeFunc, meta := pkg.Paginate(pageSize) -> db.Scopes(scopeFunc).Find(&models)
func Paginate(c *gin.Context, pageSize int) (func(db *gorm.DB) *gorm.DB, int, int, int) {
	pageGet := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageGet)

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}, page, pageSize, offset
}
