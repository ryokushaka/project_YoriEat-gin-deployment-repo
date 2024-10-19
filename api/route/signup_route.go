package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/controller"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/repository"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/usecase"
	"gorm.io/gorm"
)

// NewSignupRouter sets up the signup routes.
func NewSignupRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	tokenSecret := env.AccessTokenSecret
	tokenExpiry := env.AccessTokenExpiryHour
	refreshTokenExpiry := env.RefreshTokenExpiryHour

	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, tokenSecret, tokenExpiry, refreshTokenExpiry),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
