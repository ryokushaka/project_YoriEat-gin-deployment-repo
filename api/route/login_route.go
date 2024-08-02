package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/controller"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/repository"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/usecase"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)
	tokenSecret := env.AccessTokenSecret
	tokenExpiry := env.AccessTokenExpiryHour
	refreshTokenExpiry := env.RefreshTokenExpiryHour
	loginUsecase := usecase.NewLoginUsecase(userRepo, tokenSecret, tokenExpiry, refreshTokenExpiry)
	loginController := controller.NewLoginController(loginUsecase)

	group.POST("/login", loginController.Login)
}
