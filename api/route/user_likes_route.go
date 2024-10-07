package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/controller"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/repository"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/usecase"
	"gorm.io/gorm"
)

func NewUserLikesRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	userLikesRepo := repository.NewUserLikesRepository(db)
	userLikesUsecase := usecase.NewUserLikesUsecase(userLikesRepo)
	userLikesController := controller.NewUserLikesController(userLikesUsecase)

	group.POST("/users/:user_id/likes", userLikesController.AddLike)
	group.DELETE("/users/:user_id/likes/:recipe_id", userLikesController.RemoveLike)
	group.GET("/users/:user_id/likes", userLikesController.FetchLikesByUserID)
}
