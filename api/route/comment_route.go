package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/controller"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/repository"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/usecase"
	"gorm.io/gorm"
)

// NewCommentRouter sets up the comment routes.
func NewCommentRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	commentRepo := repository.NewCommentRepository(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepo)
	commentController := controller.NewCommentController(commentUsecase)

	group.POST("/comments", commentController.CreateComment)
	group.GET("/comments/recipe/:recipe_id", commentController.FetchCommentsByRecipeID)
	group.GET("/comments/:id", commentController.GetCommentByID)
}
