package usecase

import (
	"context"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase creates a new instance of UserUsecase.
func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uu *userUsecase) Create(c context.Context, user *domain.User) error {
	return uu.userRepo.Create(c, user)
}

func (uu *userUsecase) Fetch(c context.Context) ([]domain.User, error) {
	return uu.userRepo.Fetch(c)
}

func (uu *userUsecase) GetByEmail(c context.Context, email string) (domain.User, error) {
	return uu.userRepo.GetByEmail(c, email)
}

func (uu *userUsecase) GetByID(c context.Context, id string) (domain.User, error) {
	return uu.userRepo.GetByID(c, id)
}

func (uu *userUsecase) Update(c context.Context, user *domain.User) error {
	return uu.userRepo.Update(c, user)
}
