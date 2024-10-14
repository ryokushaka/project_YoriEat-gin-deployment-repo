package usecase

import (
	"context"
	"errors"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/internal/tokenutil"
)

type loginUsecase struct {
	userRepo           domain.UserRepository
	tokenSecret        string
	tokenExpiry        int
	refreshTokenExpiry int
}

func NewLoginUsecase(userRepo domain.UserRepository, tokenSecret string, tokenExpiry int, refreshTokenExpiry int) domain.LoginUsecase {
	return &loginUsecase{
		userRepo:           userRepo,
		tokenSecret:        tokenSecret,
		tokenExpiry:        tokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	user, err := lu.userRepo.GetByEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}
	if user.ID == 0 {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
