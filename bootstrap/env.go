package bootstrap

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// Env holds the environment variables for the application.
type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBSSLMode              string `mapstructure:"DB_SSLMODE"`
	DBPassword             string `mapstructure:"DB_PASSWORD"`
	AWSRegion              string `mapstructure:"AWS_REGION"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

// NewEnv initializes and returns a new Env instance.
func NewEnv() (*Env, error) {
	env := &Env{}

	if os.Getenv("APP_ENV") == "development" {
		viper.SetConfigFile(".env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("Warning: Error reading .env file, falling back to environment variables")
		}
	}

	// 환경 변수 또는 .env 파일에서 값을 읽어옴
	env.AppEnv = getEnv("APP_ENV", "release")
	env.ServerAddress = getEnv("SERVER_ADDRESS", ":8080")
	env.ContextTimeout = getEnvAsInt("CONTEXT_TIMEOUT", 2)
	env.DBHost = getEnv("DB_HOST", "")
	env.DBUser = getEnv("DB_USER", "")
	env.DBName = getEnv("DB_NAME", "")
	env.DBPort = getEnv("DB_PORT", "")
	env.DBSSLMode = getEnv("DB_SSLMODE", "disable")
	env.DBPassword = getEnv("DB_PASSWORD", "")
	env.AWSRegion = getEnv("AWS_REGION", "")
	env.AccessTokenExpiryHour = getEnvAsInt("ACCESS_TOKEN_EXPIRY_HOUR", 2)
	env.RefreshTokenExpiryHour = getEnvAsInt("REFRESH_TOKEN_EXPIRY_HOUR", 168)
	env.AccessTokenSecret = getEnv("ACCESS_TOKEN_SECRET", "")
	env.RefreshTokenSecret = getEnv("REFRESH_TOKEN_SECRET", "")

	requiredEnvVars := []string{
		"DB_HOST", "DB_USER", "DB_NAME", "DB_PORT",
		"DB_SSLMODE", "DB_PASSWORD", "AWS_REGION",
		"ACCESS_TOKEN_SECRET", "REFRESH_TOKEN_SECRET",
	}

	for _, v := range requiredEnvVars {
		if getEnv(v, "") == "" {
			return nil, fmt.Errorf("required environment variable not set: %s", v)
		}
	}

	if env.AppEnv == "debug" {
		log.Println("The App is running in development env")
	}

	log.Printf("APP_ENV is set to: %s", env.AppEnv)

	return env, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return viper.GetString(key)
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
