package main

import (
	"log"

	"go-echo-hexagonal/config"
	"go-echo-hexagonal/internal/core/services"
	"go-echo-hexagonal/internal/handlers"
	"go-echo-hexagonal/internal/middlewares"
	"go-echo-hexagonal/internal/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load application configurations
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Set up database connection
	db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	// Set up the layers
	userRepo := repositories.NewUserRepo(db)
	userSrv := services.NewUserSrv(userRepo, cfg.JWTSecret)
	userHdl := handlers.NewUserHdl(userSrv)
	authHdl := handlers.NewAuthHdl(userSrv)

	// Set up Echo server
	e := echo.New()

	// Routes
	e.POST("/login", authHdl.Login)
	e.POST("/users", userHdl.CreateUser)

	// Restricted routes
	r := e.Group("")
	r.Use(middlewares.Auth(cfg.JWTSecret, "admin"))
	r.GET("/users", userHdl.ListUsers)
	r.GET("/users/:id", userHdl.GetUser)


	// Start server
	if err := e.Start(cfg.ServerPort); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
