package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) error {
	result := ur.db.WithContext(ctx).Create(user)
	return result.Error
}

func (ur *userRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	result := ur.db.WithContext(ctx).Find(&users)
	return users, result.Error
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	result := ur.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, result.Error
	}
	return user, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	result := ur.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, result.Error
	}
	return user, nil
}

func (ur *userRepository) Update(ctx context.Context, user *domain.User) error {
	result := ur.db.WithContext(ctx).Save(user)
	return result.Error
}
