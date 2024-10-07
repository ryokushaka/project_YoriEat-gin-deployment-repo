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

func NewRecipeRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	recipeRepo := repository.NewRecipeRepository(db)
	recipeUsecase := usecase.NewRecipeUsecase(recipeRepo)
	recipeController := controller.NewRecipeController(recipeUsecase)

	group.POST("/recipes", recipeController.CreateRecipe)
	group.GET("/recipes", recipeController.FetchRecipes)
	group.GET("/recipes/:id", recipeController.GetRecipeByID)
	group.PUT("/recipes/:id", recipeController.UpdateRecipe)
	group.DELETE("/recipes/:id", recipeController.DeleteRecipe)
}
