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

func NewScriptRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	scriptRepo := repository.NewScriptRepository(db)
	scriptUsecase := usecase.NewScriptUsecase(scriptRepo)
	scriptController := controller.NewScriptController(scriptUsecase)

	group.POST("/scripts", scriptController.CreateScript)
	group.GET("/scripts", scriptController.FetchScripts)
	group.GET("/scripts/:id", scriptController.GetScriptByID)
	group.POST("/scripts/:script_id/recipes", scriptController.AddRecipeToScript)
	group.DELETE("/scripts/:script_id/recipes/:recipe_id", scriptController.RemoveRecipeFromScript)
}
