package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type ServerConfig struct {
	DeviceMenagerUrl string `env:"DEVICE_MANAGER_URL, required"`
}

type DatabaseConfig struct {
	Url      string `env:"DATABASE_URL, required"`
	Token     string `env:"DATABASE_TOKEN, required"`
	Org     string `env:"DATABASE_ORG, required"`
}

type BrokerConfig struct {
	TopicPattern string `env:"TOPIC_PATTERN, required"`
	BrokerUrl string `env:"BROKER_URL, required"`
}

type Config struct{
	Server   ServerConfig   
	Database DatabaseConfig
	Broker BrokerConfig
}

func loadEnv() {
	if err := godotenv.Load("../.env"); err !=nil {
		log.Fatal("Error loading .env")
	}
}

func LoadConfig() (*Config, error) {
	deploymentVariant := os.Getenv("DR_DEPLOYMENT_VARIANT")
	if deploymentVariant == "local" {
		loadEnv()
	}

	ctx := context.Background()

	config := new(Config)

	if err := envconfig.Process(ctx, config); err != nil {
		log.Fatal(err)
	  }

	return config, nil
}