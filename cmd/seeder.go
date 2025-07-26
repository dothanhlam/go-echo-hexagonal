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

	seedUsers(db)
}

func seedUsers(db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("could not hash password: %v", err)
	}

	users := []domain.User{
		{Email: "user1@example.com", Password: string(hashedPassword)},
		{Email: "user2@example.com", Password: string(hashedPassword)},
	}

	for _, user := range users {
		db.FirstOrCreate(&user, domain.User{Email: user.Email})
	}

	log.Println("Successfully seeded users")
}