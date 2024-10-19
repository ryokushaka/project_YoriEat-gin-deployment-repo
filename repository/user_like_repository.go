package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type userLikesRepository struct {
	db *gorm.DB
}

// NewUserLikesRepository creates a new instance of UserLikesRepository.
func NewUserLikesRepository(db *gorm.DB) domain.UserLikesRepository {
	return &userLikesRepository{
		db: db,
	}
}

func (ulr *userLikesRepository) AddLike(ctx context.Context, userID, recipeID int) error {
	like := domain.UserLikes{UserID: userID, RecipeID: recipeID}
	result := ulr.db.WithContext(ctx).Create(&like)
	return result.Error
}

func (ulr *userLikesRepository) RemoveLike(ctx context.Context, userID, recipeID int) error {
	result := ulr.db.WithContext(ctx).Where("user_id = ? AND recipe_id = ?", userID, recipeID).Delete(&domain.UserLikes{})
	return result.Error
}

func (ulr *userLikesRepository) FetchLikesByUserID(ctx context.Context, userID int) ([]domain.UserLikes, error) {
	var likes []domain.UserLikes
	result := ulr.db.WithContext(ctx).Where("user_id = ?", userID).Find(&likes)
	return likes, result.Error
}
