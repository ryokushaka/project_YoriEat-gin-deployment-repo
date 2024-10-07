package route

import (
	"time"

	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/middleware"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, sqlxDB *sqlx.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, sqlxDB, publicRouter)
	NewLoginRouter(env, timeout, sqlxDB, publicRouter)
	NewRefreshTokenRouter(env, timeout, sqlxDB, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewCategoryRouter(env, timeout, db, protectedRouter)
	NewRecipeRouter(env, timeout, db, protectedRouter)
	NewScriptRouter(env, timeout, db, protectedRouter)
	NewCommentRouter(env, timeout, db, protectedRouter)
	NewUserLikesRouter(env, timeout, db, protectedRouter)
	NewUserRouter(env, timeout, db, protectedRouter)

	//	Swagger endpoint
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
