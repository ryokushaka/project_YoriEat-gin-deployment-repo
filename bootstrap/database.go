package bootstrap

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/postgres"
)

func NewPostgresDatabase(env *Env) *postgres.Client {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName,
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	return &postgres.Client{DB: db}
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