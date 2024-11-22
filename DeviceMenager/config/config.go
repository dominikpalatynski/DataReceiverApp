package config

import (
	"ConfigApp/logging"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type ServerConfig struct {
	Port string    `env:"SERVER_PORT, required"`
	AuthCookieName string `env:"AUTH_COOKIE_NAME, required"`
}

type DatabaseConfig struct {
	Url      string `env:"DATABASE_URL, required"`
	Key     string `env:"DATABASE_KEY, required"`
}

type CacheConfig struct {
	Url      string `env:"CACHE_URL, required"`
	Password     string `env:"CACHE_PASSWORD, required"`
}

type Config struct{
	Server   ServerConfig
	Database DatabaseConfig
	Cache CacheConfig
}

func loadEnv() {
	if err := godotenv.Load(".env"); err !=nil {
		log.Fatal("Error loading .env")
	}
}

func LoadConfig() (*Config, error) {
	deploymentVariant := os.Getenv("DM_DEPLOYMENT_VARIANT")
	if deploymentVariant == "local" {
		loadEnv()
	}

	ctx := context.Background()

	config := new(Config)

	if err := envconfig.Process(ctx, config); err != nil {
		logging.Log.Fatalf("Cannot load configuration: %v", err)
	  }

	logging.Log.Info("Configuration loaded: %v", config)

	return config, nil
}
