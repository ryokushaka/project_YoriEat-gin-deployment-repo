package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) domain.RecipeRepository {
	return &recipeRepository{
		db: db,
	}
}

func (rr *recipeRepository) Create(ctx context.Context, recipe *domain.Recipe) error {
	result := rr.db.WithContext(ctx).Create(recipe)
	return result.Error
}

func (rr *recipeRepository) Fetch(ctx context.Context) ([]domain.Recipe, error) {
	var recipes []domain.Recipe
	result := rr.db.WithContext(ctx).Find(&recipes)
	if result.Error != nil {
		return nil, result.Error
	}
	return recipes, nil
}

func (rr *recipeRepository) GetByID(ctx context.Context, id string) (domain.Recipe, error) {
	var recipe domain.Recipe
	result := rr.db.WithContext(ctx).First(&recipe, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return recipe, nil
		}
		return recipe, result.Error
	}
	return recipe, nil
}

func (rr *recipeRepository) Update(ctx context.Context, recipe *domain.Recipe) error {
	result := rr.db.WithContext(ctx).Save(recipe)
	return result.Error
}

func (rr *recipeRepository) Delete(ctx context.Context, id string) error {
	result := rr.db.WithContext(ctx).Delete(&domain.Recipe{}, id)
	return result.Error
}
