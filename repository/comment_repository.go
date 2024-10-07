package repository

import (
	"context"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	result := cr.db.WithContext(ctx).Create(comment)
	return result.Error
}

func (cr *commentRepository) FetchByRecipeID(ctx context.Context, recipeID int) ([]domain.Comment, error) {
	var comments []domain.Comment
	result := cr.db.WithContext(ctx).Where("recipe_id = ?", recipeID).Find(&comments)
	return comments, result.Error
}

func (cr *commentRepository) GetByID(ctx context.Context, id string) (domain.Comment, error) {
	var comment domain.Comment
	result := cr.db.WithContext(ctx).First(&comment, id)
	return comment, result.Error
}
