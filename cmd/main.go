package main

import (
	"log/slog"
	"os"
	"pulse-auth/internal/config"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.New()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("Start server", slog.String("address", cfg.Address))
	log.Debug("Debug mode enable")
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envDev:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	}
	return logger
}
