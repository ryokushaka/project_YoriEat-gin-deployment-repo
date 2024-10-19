package domain

import (
	"context"
)

// UserLikes represents a user's like on a recipe.
type UserLikes struct {
	UserID   int `json:"user_id"`
	RecipeID int `json:"recipe_id"`
}

// UserLikesUsecase represents the user likes use case interface.
type UserLikesUsecase interface {
	AddLike(c context.Context, userID, recipeID int) error
	RemoveLike(c context.Context, userID, recipeID int) error
	FetchLikesByUserID(c context.Context, userID int) ([]UserLikes, error)
}

// UserLikesRepository represents the user likes repository interface.
type UserLikesRepository interface {
	AddLike(c context.Context, userID, recipeID int) error
	RemoveLike(c context.Context, userID, recipeID int) error
	FetchLikesByUserID(c context.Context, userID int) ([]UserLikes, error)
}
