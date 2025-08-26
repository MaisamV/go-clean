package config

import (
	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// Load loads configuration from environment variables and config files
func Load() (*Config, error) {
	viper.SetDefault("server.port", "8080")

	// Set environment variable prefix
	viper.SetEnvPrefix("GO_CLEAN")
	viper.AutomaticEnv()

	// Try to read from config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Reading config file is optional
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}