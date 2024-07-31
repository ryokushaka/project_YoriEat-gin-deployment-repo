package domain

import (
	"context"
)

type Comment struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	UserID    int    `json:"user_id"`
	RecipeID  int    `json:"recipe_id"`
	CommentID *int   `json:"comment_id"`
}

type CommentUsecase interface {
	Create(c context.Context, comment *Comment) error
	FetchByRecipeID(c context.Context, recipeID int) ([]Comment, error)
	GetByID(c context.Context, id string) (Comment, error)
}

type CommentRepository interface {
	Create(c context.Context, comment *Comment) error
	FetchByRecipeID(c context.Context, recipeID int) ([]Comment, error)
	GetByID(c context.Context, id string) (Comment, error)
}
