package main

import (
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/route"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
	_ "github.com/ryokushaka/project_YoriEat-gin-deployment-repo/cmd/docs"
)

func main() {
	// 애플리케이션 초기화
	app := bootstrap.App()

	// 데이터베이스 연결 확인
	err := app.Postgres.DB.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	env := app.Env

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()

	router.Use(gin.Logger())      // 요청에 대한 로깅을 수행
	router.Use(gin.Recovery())    // Panic이 발생했을 때 복구하고 500 에러를 반환
	router.Use(middleware.CORS()) // CORS 미들웨어 추가

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.Setup(env, timeout, app.Postgres.DB, router)

	router.Run(env.ServerAddress)
}
