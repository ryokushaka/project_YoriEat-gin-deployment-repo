package usecase

import (
	"context"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type userLikesUsecase struct {
	userLikesRepo domain.UserLikesRepository
}

// NewUserLikesUsecase creates a new instance of UserLikesUsecase.
func NewUserLikesUsecase(userLikesRepo domain.UserLikesRepository) domain.UserLikesUsecase {
	return &userLikesUsecase{
		userLikesRepo: userLikesRepo,
	}
}

func (ulu *userLikesUsecase) AddLike(c context.Context, userID, recipeID int) error {
	return ulu.userLikesRepo.AddLike(c, userID, recipeID)
}

func (ulu *userLikesUsecase) RemoveLike(c context.Context, userID, recipeID int) error {
	return ulu.userLikesRepo.RemoveLike(c, userID, recipeID)
}

func (ulu *userLikesUsecase) FetchLikesByUserID(c context.Context, userID int) ([]domain.UserLikes, error) {
	return ulu.userLikesRepo.FetchLikesByUserID(c, userID)
}
