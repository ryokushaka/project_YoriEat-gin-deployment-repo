package usecase

import (
	"context"
	"errors"

	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/internal/tokenutil"
)

type refreshTokenUsecase struct {
	userRepo           domain.UserRepository
	tokenSecret        string
	tokenExpiry        int
	refreshTokenExpiry int
}

// NewRefreshTokenUsecase creates a new instance of RefreshTokenUsecase.
func NewRefreshTokenUsecase(userRepo domain.UserRepository, tokenSecret string, tokenExpiry int, refreshTokenExpiry int) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepo:           userRepo,
		tokenSecret:        tokenSecret,
		tokenExpiry:        tokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	user, err := rtu.userRepo.GetByID(c, id)
	if err != nil {
		return domain.User{}, err
	}
	if user.ID == 0 {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
