package usecase

import (
	"context"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type commentUsecase struct {
	commentRepo domain.CommentRepository
}

// NewCommentUsecase creates a new instance of CommentUsecase.
func NewCommentUsecase(commentRepo domain.CommentRepository) domain.CommentUsecase {
	return &commentUsecase{
		commentRepo: commentRepo,
	}
}

func (cu *commentUsecase) Create(c context.Context, comment *domain.Comment) error {
	return cu.commentRepo.Create(c, comment)
}

func (cu *commentUsecase) FetchByRecipeID(c context.Context, recipeID int) ([]domain.Comment, error) {
	return cu.commentRepo.FetchByRecipeID(c, recipeID)
}

func (cu *commentUsecase) GetByID(c context.Context, id string) (domain.Comment, error) {
	return cu.commentRepo.GetByID(c, id)
}
