package logger

import (
	"log"

	"go.uber.org/zap"
)

func SetupLogger(env string) *zap.Logger {
	var logger *zap.Logger
	var err error

	switch env {
	case "dev", "local":
		logger, err = zap.NewDevelopment()
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	zap.ReplaceGlobals(logger)

	return logger
}
