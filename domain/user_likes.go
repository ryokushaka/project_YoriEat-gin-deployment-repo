package domain

import (
	"context"
)

type UserLikes struct {
	UserID   int `json:"user_id"`
	RecipeID int `json:"recipe_id"`
}

type UserLikesUsecase interface {
	AddLike(c context.Context, userID, recipeID int) error
	RemoveLike(c context.Context, userID, recipeID int) error
	FetchLikesByUserID(c context.Context, userID int) ([]UserLikes, error)
}

type UserLikesRepository interface {
	AddLike(c context.Context, userID, recipeID int) error
	RemoveLike(c context.Context, userID, recipeID int) error
	FetchLikesByUserID(c context.Context, userID int) ([]UserLikes, error)
}
