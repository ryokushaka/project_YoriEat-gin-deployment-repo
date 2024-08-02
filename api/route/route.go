package route

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/middleware"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	//NewUserRouter(env, timeout, db, protectedRouter)

	//	Swagger endpoint
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
