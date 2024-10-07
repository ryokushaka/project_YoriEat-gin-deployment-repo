package bootstrap

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

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

func NewEnv() (*Env, error) {
	env := &Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = viper.Unmarshal(env)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	requiredEnvVars := []string{
		"DB_HOST", "DB_USER", "DB_NAME", "DB_PORT",
		"DB_SSLMODE", "DB_PASSWORD", "AWS_REGION",
		"ACCESS_TOKEN_SECRET", "REFRESH_TOKEN_SECRET",
	}

	for _, v := range requiredEnvVars {
		if viper.GetString(v) == "" {
			return nil, fmt.Errorf("required environment variable not set: %s", v)
		}
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return env, nil
}
