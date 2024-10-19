package domain

import (
	"context"
)

// Comment represents a comment entity.
type Comment struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	UserID    int    `json:"user_id"`
	RecipeID  int    `json:"recipe_id"`
	CommentID *int   `json:"comment_id"`
}

// CommentUsecase represents the comment use case interface.
type CommentUsecase interface {
	Create(c context.Context, comment *Comment) error
	FetchByRecipeID(c context.Context, recipeID int) ([]Comment, error)
	GetByID(c context.Context, id string) (Comment, error)
}

// CommentRepository represents the comment repository interface.
type CommentRepository interface {
	Create(c context.Context, comment *Comment) error
	FetchByRecipeID(c context.Context, recipeID int) ([]Comment, error)
	GetByID(c context.Context, id string) (Comment, error)
}
