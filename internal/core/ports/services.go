package ports

import (
	"context"

	"go-echo-hexagonal/internal/core/domain"
)

// UserService is an interface for business logic related to users.
type UserService interface {
	CreateUser(ctx context.Context, email, password string) (*domain.User, error)
	GetUser(ctx context.Context, id uint) (*domain.User, error)
	ListUsers(ctx context.Context, page, limit int) (*domain.Paginator, error)
	Login(ctx context.Context, email, password string) (string, error)
}
