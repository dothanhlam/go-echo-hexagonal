package services

import (
	"context"
	"time"

	"go-echo-hexagonal/internal/core/domain"
	"go-echo-hexagonal/internal/core/ports"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// UserSrv implements the UserService interface.
type UserSrv struct {
	repo       ports.UserRepository
	jwtSecret  string
}

// NewUserSrv creates a new UserService.
func NewUserSrv(repo ports.UserRepository, jwtSecret string) *UserSrv {
	return &UserSrv{repo: repo, jwtSecret: jwtSecret}
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
		Role:      "user",
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

// ListUsers retrieves a paginated list of users.
func (s *UserSrv) ListUsers(ctx context.Context, page, limit int) (*domain.Paginator, error) {
	return s.repo.FindAll(ctx, page, limit)
}

// Login authenticates a user and returns a JWT.
func (s *UserSrv) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}
