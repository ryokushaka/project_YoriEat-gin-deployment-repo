package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/middleware"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/route"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	_ "github.com/ryokushaka/project_YoriEat-gin-deployment-repo/cmd/docs"
)

func main() {
	app, err := initializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer app.CloseDBConnection()

	if err := checkDatabaseConnection(app); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	router := setupRouter(app)
	router.Run(app.Env.ServerAddress)
}

func initializeApp() (*bootstrap.Application, error) {
	return bootstrap.App()
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

func setupRouter(app *bootstrap.Application) *gin.Engine {
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	router.GET("/health", healthCheckHandler(app))

	route.Setup(env, timeout, app.DB, router)
	return router
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
