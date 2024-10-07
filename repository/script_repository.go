package repository

import (
	"context"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
	"gorm.io/gorm"
)

type scriptRepository struct {
	db *gorm.DB
}

func NewScriptRepository(db *gorm.DB) domain.ScriptRepository {
	return &scriptRepository{
		db: db,
	}
}

func (sr *scriptRepository) Create(ctx context.Context, script *domain.Script) error {
	result := sr.db.WithContext(ctx).Create(script)
	return result.Error
}

func (sr *scriptRepository) Fetch(ctx context.Context) ([]domain.Script, error) {
	var scripts []domain.Script
	result := sr.db.WithContext(ctx).Find(&scripts)
	return scripts, result.Error
}

func (sr *scriptRepository) GetByID(ctx context.Context, id string) (domain.Script, error) {
	var script domain.Script
	result := sr.db.WithContext(ctx).First(&script, id)
	return script, result.Error
}

func (sr *scriptRepository) AddRecipeToScript(ctx context.Context, scriptID, recipeID int) error {
	scriptRecipe := domain.ScriptRecipe{ScriptID: scriptID, RecipeID: recipeID}
	result := sr.db.WithContext(ctx).Create(&scriptRecipe)
	return result.Error
}

func (sr *scriptRepository) RemoveRecipeFromScript(ctx context.Context, scriptID, recipeID int) error {
	result := sr.db.WithContext(ctx).Where("script_id = ? AND recipe_id = ?", scriptID, recipeID).Delete(&domain.ScriptRecipe{})
	return result.Error
}
