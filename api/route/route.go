// Package route provides the routing setup for the application.
package route

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/middleware"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
)

// Setup initializes the routing for the application.
func Setup(app *bootstrap.Application, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(app.Env, db, publicRouter)
	NewLoginRouter(app.Env, db, publicRouter)
	NewRefreshTokenRouter(app.Env, db, publicRouter)
	publicRouter.GET("/health", healthCheckHandler(app))

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(app.Env.AccessTokenSecret))
	// All Private APIs
	NewCategoryRouter(app.Env, db, protectedRouter)
	NewRecipeRouter(app.Env, db, protectedRouter)
	NewScriptRouter(app.Env, db, protectedRouter)
	NewCommentRouter(app.Env, db, protectedRouter)
	NewUserLikesRouter(app.Env, db, protectedRouter)
	NewUserRouter(app.Env, db, protectedRouter)

	// Swagger endpoint
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func checkDatabaseConnection(app *bootstrap.Application) error {
	sqlDB, err := app.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("sqlx database connection failed: %w", err)
	}

	return nil
}

func healthCheckHandler(app *bootstrap.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := checkDatabaseConnection(app)
		if err != nil {
			c.JSON(500, gin.H{"status": "DOWN", "reason": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "UP"})
	}
}
