package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	// Environment variable keys
	EnvPort   = "PORT"
	EnvAppEnv = "APP_ENV"

	// Default values
	DefaultPort   = "8080"
	DefaultAppEnv = "production"
)

// Config holds the application configuration
type Config struct {
	Port   string
	AppEnv string
}

// Load reads configuration from the environment and loads the .env file in the project root if available
// Sets default values if applicable
func Load(logger *zap.Logger) (*Config, error) {
	// Load environment variables from .env if available.
	err := godotenv.Load()
	if err != nil {
		logger.Warn("no .env file available, using default values")
	}

	port, ok := os.LookupEnv(EnvPort)
	if !ok {
		port = DefaultPort
	}

	appEnv, ok := os.LookupEnv(EnvAppEnv)
	if !ok {
		appEnv = DefaultAppEnv
	}

	cfg := &Config{
		Port:   port,
		AppEnv: appEnv,
	}

	logger.Info("Configuration loaded",
		zap.String("Port", cfg.Port),
		zap.String("AppEnv", cfg.AppEnv),
	)

	return cfg, nil
}
