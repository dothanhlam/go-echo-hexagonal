package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go-echo-hexagonal/internal/core/domain"
	"go-echo-hexagonal/internal/pkg/utils"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// UserRepo implements the UserRepository interface.
type UserRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewUserRepo creates a new UserRepository.
func NewUserRepo(db *gorm.DB, rdb *redis.Client) *UserRepo {
	return &UserRepo{db: db, rdb: rdb}
}

// Save saves a user to the database.
func (r *UserRepo) Save(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindByID finds a user by ID.
func (r *UserRepo) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	key := fmt.Sprintf("user:%d", id)

	// Try to get the user from Redis first
	val, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &user); err == nil {
			return &user, nil
		}
	}

	// If not in Redis, get from the database
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Cache the user in Redis
	jsonUser, err := json.Marshal(user)
	if err == nil {
		r.rdb.Set(ctx, key, jsonUser, time.Hour)
	}

	return &user, nil
}

// FindByEmail finds a user by email.
func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	key := fmt.Sprintf("user:%s", email)

	// Try to get the user from Redis first
	val, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &user); err == nil {
			return &user, nil
		}
	}

	// If not in Redis, get from the database
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Cache the user in Redis
	jsonUser, err := json.Marshal(user)
	if err == nil {
		r.rdb.Set(ctx, key, jsonUser, time.Hour)
	}

	return &user, nil
}

// FindAll finds all users with pagination.
func (r *UserRepo) FindAll(ctx context.Context, page, limit int) (*domain.Paginator, error) {
	var users []domain.User
	paginator, err := utils.Paginator(r.db.WithContext(ctx), page, limit, &users)
	if err != nil {
		return nil, err
	}
	return paginator, nil
}
