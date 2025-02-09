package logger

import (
	"go.uber.org/zap"
)

// NewLogger creates a new zap.Logger instance
func NewLogger(appEnv string) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	logger, err = zap.NewProduction()

	if appEnv == "development" {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	logger = logger.With(zap.String("service", "postr-backend"), zap.String("version", "1.0.0"))

	return logger, nil
}
