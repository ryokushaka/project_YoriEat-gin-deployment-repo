package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/controller"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/repository"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/usecase"
	"gorm.io/gorm"
)

// NewRecipeRouter sets up the recipe routes.
func NewRecipeRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	recipeRepo := repository.NewRecipeRepository(db)
	recipeUsecase := usecase.NewRecipeUsecase(recipeRepo)
	recipeController := controller.NewRecipeController(recipeUsecase)

	group.POST("/recipes", recipeController.CreateRecipe)
	group.GET("/recipes", recipeController.FetchRecipes)
	group.GET("/recipes/:id", recipeController.GetRecipeByID)
	group.PUT("/recipes/:id", recipeController.UpdateRecipe)
	group.DELETE("/recipes/:id", recipeController.DeleteRecipe)
}
