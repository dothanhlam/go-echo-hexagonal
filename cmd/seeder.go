package main

import (
	"log"

	"go-echo-hexagonal/config"
	"go-echo-hexagonal/internal/core/domain"

	"golang.org/x/crypto/bcrypt"
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

	// Drop the table and recreate it
	db.Migrator().DropTable(&domain.User{})
	db.AutoMigrate(&domain.User{})

	seedUsers(db)
}

func seedUsers(db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("could not hash password: %v", err)
	}

	users := []domain.User{
		{Email: "admin@example.com", Password: string(hashedPassword), Role: "admin"},
		{Email: "user@example.com", Password: string(hashedPassword), Role: "user"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	log.Println("Successfully seeded users")
}
