package main

import (
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

	if app.Env.AppEnv == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	if err := checkDatabaseConnection(app); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	router := setupRouter(app)
	if err := router.Run(app.Env.ServerAddress); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}

func initializeApp() (*bootstrap.Application, error) {
	return bootstrap.App()
}

func setupRouter(app *bootstrap.Application) *gin.Engine {
	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	route.Setup(app, timeout, app.DB, router)
	return router
}

func checkDatabaseConnection(app *bootstrap.Application) error {
	sqlDB, err := app.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
