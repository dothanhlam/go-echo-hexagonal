package repositories

import (
	"context"
	"fmt"

	"go-echo-hexagonal/internal/core/domain"

	"gorm.io/gorm"
)

// UserRepo implements the UserRepository interface.
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo creates a new UserRepository.
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Save saves a user to the database.
func (r *UserRepo) Save(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindByID finds a user by ID.
func (r *UserRepo) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}
