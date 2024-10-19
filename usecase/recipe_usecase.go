package usecase

import (
	"context"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type recipeUsecase struct {
	recipeRepo domain.RecipeRepository
}

// NewRecipeUsecase creates a new instance of RecipeUsecase.
func NewRecipeUsecase(recipeRepo domain.RecipeRepository) domain.RecipeUsecase {
	return &recipeUsecase{
		recipeRepo: recipeRepo,
	}
}

func (ru *recipeUsecase) Create(c context.Context, recipe *domain.Recipe) error {
	return ru.recipeRepo.Create(c, recipe)
}

func (ru *recipeUsecase) Fetch(c context.Context) ([]domain.Recipe, error) {
	return ru.recipeRepo.Fetch(c)
}

func (ru *recipeUsecase) GetByID(c context.Context, id string) (domain.Recipe, error) {
	return ru.recipeRepo.GetByID(c, id)
}

func (ru *recipeUsecase) Update(c context.Context, recipe *domain.Recipe) error {
	return ru.recipeRepo.Update(c, recipe)
}

func (ru *recipeUsecase) Delete(c context.Context, id string) error {
	return ru.recipeRepo.Delete(c, id)
}
