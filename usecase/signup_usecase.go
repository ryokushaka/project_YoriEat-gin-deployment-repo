package usecase

import (
	"context"
	"errors"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/internal/tokenutil"
)

type signupUsecase struct {
	userRepo           domain.UserRepository
	tokenSecret        string
	tokenExpiry        int
	refreshTokenExpiry int
}

func NewSignupUsecase(userRepo domain.UserRepository, tokenSecret string, tokenExpiry int, refreshTokenExpiry int) domain.SignupUsecase {
	return &signupUsecase{
		userRepo:           userRepo,
		tokenSecret:        tokenSecret,
		tokenExpiry:        tokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	existingUser, err := su.userRepo.GetByEmail(c, user.Email)
	if err != nil {
		return err
	}
	if existingUser.ID != 0 {
		return errors.New("email already in use")
	}
	return su.userRepo.Create(c, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	return su.userRepo.GetByEmail(c, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
