package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/controller"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/repository"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/usecase"
	"gorm.io/gorm"
)

// NewCategoryRouter sets up the category routes.
func NewCategoryRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryController := controller.NewCategoryController(categoryUsecase)

	group.POST("/categories", categoryController.CreateCategory)
	group.GET("/categories", categoryController.FetchCategories)
	group.GET("/categories/:id", categoryController.GetCategoryByID)
}
