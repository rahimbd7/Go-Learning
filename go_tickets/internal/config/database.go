package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	fmt.Println("Database connection successful")
	return db
}
