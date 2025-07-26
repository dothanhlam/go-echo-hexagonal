package services

import (
	"context"
	"time"

	"go-echo-hexagonal/internal/core/domain"
	"go-echo-hexagonal/internal/core/ports"

	"golang.org/x/crypto/bcrypt"
)

// UserSrv implements the UserService interface.
type UserSrv struct {
	repo ports.UserRepository
}

// NewUserSrv creates a new UserService.
func NewUserSrv(repo ports.UserRepository) *UserSrv {
	return &UserSrv{repo: repo}
}

// CreateUser creates a new user.
func (s *UserSrv) CreateUser(ctx context.Context, email, password string) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID.
func (s *UserSrv) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.FindByID(ctx, id)
}
