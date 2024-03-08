package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"pulse-auth/cmd/pulse/application"
	"pulse-auth/cmd/pulse/config"
	"pulse-auth/internal/utils"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.New()

	logger, err := constructLogger(cfg.Logger)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	app := application.New(cfg, logger)
	if err = app.Run(); err != nil {
		logger.Sugar().Fatalf("application stopped with error: %v", err)
	} else {
		logger.Info("application stopped")
	}
}

func constructLogger(cfg config.LoggerConfig) (*zap.Logger, error) {
	var logger *zap.Logger

	var err error

	switch cfg.Environment {
	case utils.Development:
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, fmt.Errorf("new development logger")
		}
	case utils.Production:
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("new production logger")
		}
	default:
		return nil, fmt.Errorf("unexpected environment for logger: %w", err)
	}

	defer func() { _ = logger.Sync() }()

	return logger, nil
}
