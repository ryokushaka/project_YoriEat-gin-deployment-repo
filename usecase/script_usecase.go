package usecase

import (
	"context"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type scriptUsecase struct {
	scriptRepo domain.ScriptRepository
}

func NewScriptUsecase(scriptRepo domain.ScriptRepository) domain.ScriptUsecase {
	return &scriptUsecase{
		scriptRepo: scriptRepo,
	}
}

func (su *scriptUsecase) Create(c context.Context, script *domain.Script) error {
	return su.scriptRepo.Create(c, script)
}

func (su *scriptUsecase) Fetch(c context.Context) ([]domain.Script, error) {
	return su.scriptRepo.Fetch(c)
}

func (su *scriptUsecase) GetByID(c context.Context, id string) (domain.Script, error) {
	return su.scriptRepo.GetByID(c, id)
}

func (su *scriptUsecase) AddRecipeToScript(c context.Context, scriptID, recipeID int) error {
	return su.scriptRepo.AddRecipeToScript(c, scriptID, recipeID)
}

func (su *scriptUsecase) RemoveRecipeFromScript(c context.Context, scriptID, recipeID int) error {
	return su.scriptRepo.RemoveRecipeFromScript(c, scriptID, recipeID)
}
