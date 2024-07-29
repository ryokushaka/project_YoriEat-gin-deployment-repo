package bootstrap

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/postgres"
)

func NewPostgresDatabase(env *Env) *postgres.Client {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	client := &postgres.Client{
		DB: db,
	}

	return client
}

func ClosePostgresDBConnection(client *postgres.Client) {
	if client == nil || client.DB == nil {
		return
	}

	err := client.DB.Close()
	if err != nil {
		log.Fatal("Failed to close PostgreSQL connection:", err)
	}

	log.Println("Connection to PostgreSQL closed.")
}