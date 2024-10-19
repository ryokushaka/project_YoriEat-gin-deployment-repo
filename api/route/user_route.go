package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/controller"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/repository"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/usecase"
	"gorm.io/gorm"
)

// NewUserRouter sets up the user routes.
func NewUserRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	group.POST("/users", userController.CreateUser)
	group.GET("/users", userController.FetchUsers)
	group.GET("/users/email/:email", userController.GetUserByEmail)
	group.GET("/users/:user_id", userController.GetUserByID)
	group.PUT("/users/:user_id", userController.UpdateUser)
}
