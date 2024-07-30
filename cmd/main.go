package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	route "github.com/ryokushaka/project_YoriEat-gin-deployment-repo/api/route"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/bootstrap"
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

	route.Setup(env, timeout, app.Postgres.DB, router)

	router.Run(env.ServerAddress)
}