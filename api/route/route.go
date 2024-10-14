package route

import (
	"time"

	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/middleware"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, db, publicRouter)
	NewLoginRouter(env, db, publicRouter)
	NewRefreshTokenRouter(env, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewCategoryRouter(env, db, protectedRouter)
	NewRecipeRouter(env, db, protectedRouter)
	NewScriptRouter(env, db, protectedRouter)
	NewCommentRouter(env, db, protectedRouter)
	NewUserLikesRouter(env, db, protectedRouter)
	NewUserRouter(env, db, protectedRouter)

	//	Swagger endpoint
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
