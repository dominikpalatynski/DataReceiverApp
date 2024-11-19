package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	DeviceMenagerUrl string `mapstructure:"device_menager_url"`
}

type DatabaseConfig struct {
	Url      string `mapstructure:"database_url"`
	Token     string `mapstructure:"database_token"`
	Org     string `mapstructure:"database_org"`
}

type BrokerConfig struct {
	TopicPattern string `mapstructure:"topic_pattern"`
	BrokerUrl string `mapstructure:"broker_url"`
}

type Config struct{
	Server   ServerConfig   `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
	Broker BrokerConfig `mapstructure:",squash"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("błąd ładowania pliku .env: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("błąd mapowania konfiguracji: %w", err)
	}

	return &config, nil
}
