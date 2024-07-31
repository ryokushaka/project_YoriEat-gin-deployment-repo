package domain

import (
	"context"
)

type Script struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

type ScriptRecipe struct {
	ScriptID int `json:"script_id"`
	RecipeID int `json:"recipe_id"`
}

type ScriptUsecase interface {
	Create(c context.Context, script *Script) error
	Fetch(c context.Context) ([]Script, error)
	GetByID(c context.Context, id string) (Script, error)
	AddRecipeToScript(c context.Context, scriptID, recipeID int) error
	RemoveRecipeFromScript(c context.Context, scriptID, recipeID int) error
}

type ScriptRepository interface {
	Create(c context.Context, script *Script) error
	Fetch(c context.Context) ([]Script, error)
	GetByID(c context.Context, id string) (Script, error)
	AddRecipeToScript(c context.Context, scriptID, recipeID int) error
	RemoveRecipeFromScript(c context.Context, scriptID, recipeID int) error
}
