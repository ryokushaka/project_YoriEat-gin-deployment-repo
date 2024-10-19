package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new instance of CategoryRepository.
func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (cr *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	result := cr.db.WithContext(ctx).Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *categoryRepository) Fetch(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	result := cr.db.WithContext(ctx).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (cr *categoryRepository) GetByID(ctx context.Context, id string) (domain.Category, error) {
	var category domain.Category
	result := cr.db.WithContext(ctx).First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return category, nil
		}
		return category, result.Error
	}
	return category, nil
}
