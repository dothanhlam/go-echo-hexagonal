package main

import (
	"log"

	"go-echo-hexagonal/config"
	"go-echo-hexagonal/internal/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	db.AutoMigrate(&domain.User{})

	log.Println("Successfully migrated database")
}
