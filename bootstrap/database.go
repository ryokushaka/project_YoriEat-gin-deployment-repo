package bootstrap

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db     *gorm.DB
	sqlxDB *sqlx.DB
)

func initializeDB(env *Env) error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Seoul", env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort, env.DBSSLMode)

	db, err = connectWithRetry(dsn, 5, 5*time.Second)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// sqlx 연결 설정
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	sqlxDB = sqlx.NewDb(sqlDB, "postgres")

	// Connection Pool 설정
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

// GetDB returns the gorm.DB instance.
func GetDB() *gorm.DB {
	return db
}

// GetSqlxDB returns the sqlx.DB instance.
func GetSqlxDB() *sqlx.DB {
	return sqlxDB
}

// NewPostgresDatabase initializes and returns the gorm.DB and sqlx.DB instances.
func NewPostgresDatabase() (*gorm.DB, *sqlx.DB, error) {
	env, err := NewEnv()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	err = initializeDB(env)
	if err != nil {
		return nil, nil, err
	}

	return db, sqlxDB, nil
}

// ClosePostgresDBConnection closes the database connections.
func ClosePostgresDBConnection() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting DB instance: %v", err)
			return
		}

		err = sqlDB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
			return
		}

		log.Println("Database connection closed successfully")
	}

	if sqlxDB != nil {
		err := sqlxDB.Close()
		if err != nil {
			log.Printf("Error closing sqlx database connection: %v", err)
		}
	}
}

func connectWithRetry(dsn string, maxRetries int, retryInterval time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			return db, nil
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		log.Printf("Error details: %+v", err)
		time.Sleep(retryInterval)
	}
	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}
