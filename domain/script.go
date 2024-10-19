package domain

import (
	"context"
)

// Script represents a script entity.
type Script struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

// ScriptRecipe represents the relationship between a script and a recipe.
type ScriptRecipe struct {
	ScriptID int `json:"script_id"`
	RecipeID int `json:"recipe_id"`
}

// ScriptUsecase represents the script use case interface.
type ScriptUsecase interface {
	Create(c context.Context, script *Script) error
	Fetch(c context.Context) ([]Script, error)
	GetByID(c context.Context, id string) (Script, error)
	AddRecipeToScript(c context.Context, scriptID, recipeID int) error
	RemoveRecipeFromScript(c context.Context, scriptID, recipeID int) error
}

// ScriptRepository represents the script repository interface.
type ScriptRepository interface {
	Create(c context.Context, script *Script) error
	Fetch(c context.Context) ([]Script, error)
	GetByID(c context.Context, id string) (Script, error)
	AddRecipeToScript(c context.Context, scriptID, recipeID int) error
	RemoveRecipeFromScript(c context.Context, scriptID, recipeID int) error
}
