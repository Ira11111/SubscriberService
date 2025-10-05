package logger

import (
	"log/slog"
	"os"
)

const (
	envlocal       = "local"
	envdevelopment = "dev"
	envproduction  = "prod"
)

func InitLogger(envType string) *slog.Logger {
	var logger *slog.Logger

	switch envType {
	case envlocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envdevelopment:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envproduction:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
