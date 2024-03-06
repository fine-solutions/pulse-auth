package main

import (
	"context"
	"log/slog"
	"os"
	"pulse-auth/internal/config"
	"pulse-auth/internal/storage"
	"pulse-auth/internal/storage/postgres"
	"time"
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

	db, err := postgres.New(log, cfg)
	if err != nil {
		log.Error("Can't init database", err)
		os.Exit(1)
	}
	log.Info("Success connect to database")

	db.CreateUser(context.Background(), &storage.User{
		Username:     "Nikita",
		PasswordHash: "asdasd",
		Email:        "some@email",
		CreatedAt:    time.Now(),
	})
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
