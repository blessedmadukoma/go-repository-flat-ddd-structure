package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	SIGNINGKEY string `mapstructure:"SIGNING_KEY"`
}

func LoadEnvConfig() (config Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load env:", err)
	}

	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBSource = os.Getenv("DB_SOURCE")

	return config
}

func CorsWhiteList() string {
	return os.Getenv("CORS_WHITELIST")
}
