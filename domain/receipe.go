package domain

import "context"

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Text        string   `json:"text"`
	Ingredient  []string `json:"ingredient"`
	Time        int      `json:"time"`
	Process     []string `json:"process"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	CategoryID  int      `json:"category_id"`
	UserID      int      `json:"user_id"`
}

type RecipeUsecase interface {
	Create(c context.Context, recipe *Recipe) error
	Fetch(c context.Context) ([]Recipe, error)
	GetByID(c context.Context, id string) (Recipe, error)
	Update(c context.Context, recipe *Recipe) error
	Delete(c context.Context, id string) error
}

type RecipeRepository interface {
	Create(c context.Context, recipe *Recipe) error
	Fetch(c context.Context) ([]Recipe, error)
	GetByID(c context.Context, id string) (Recipe, error)
	Update(c context.Context, recipe *Recipe) error
	Delete(c context.Context, id string) error
}
