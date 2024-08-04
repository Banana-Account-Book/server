package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"banana-account-book.com/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() *gorm.DB {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
			LogLevel:                  logger.Info, // Log level
		},
	)

	var sslMode string
	if config.IsProduction {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=UTC sslmode=%s", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
