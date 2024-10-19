package usecase

import (
	"context"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
}

// NewCategoryUsecase creates a new instance of CategoryUsecase.
func NewCategoryUsecase(categoryRepo domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
	}
}

func (cu *categoryUsecase) Create(c context.Context, category *domain.Category) error {
	return cu.categoryRepo.Create(c, category)
}

func (cu *categoryUsecase) Fetch(c context.Context) ([]domain.Category, error) {
	return cu.categoryRepo.Fetch(c)
}

func (cu *categoryUsecase) GetByID(c context.Context, id string) (domain.Category, error) {
	return cu.categoryRepo.GetByID(c, id)
}
