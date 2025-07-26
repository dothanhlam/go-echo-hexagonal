package ports

import (
	"context"

	"go-echo-hexagonal/internal/core/domain"
)

// UserRepository is an interface for interacting with user-related data.
type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id uint) (*domain.User, error)
}
