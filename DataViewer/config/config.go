package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string    `mapstructure:"server_port"`
	AuthCookieName string `mapstructure:"auth_cookie_name"`
}

type DatabaseConfig struct {
	Url      string `mapstructure:"database_url"`
	Key     string `mapstructure:"database_key"`
	Org string `mapstructure:"database_org"`
}

type Config struct{
	Server   ServerConfig   `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
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
