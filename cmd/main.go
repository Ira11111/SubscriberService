package main

import (
	"SubscriberService/internal/app"
	"SubscriberService/internal/config"
	l "SubscriberService/internal/logger"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoadByPath(".env")
	logger := l.InitLogger(cfg.Env)
	logger.Debug("Init logger")

	fmt.Println(cfg)
	application := app.NewApp(cfg, logger)
	logger.Debug("Init application")

	go application.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		logger.Info("Shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := application.Shutdown(ctx); err != nil {
			logger.Warn("stop app error", slog.Any("err", err.Error()))
			os.Exit(1)
		}
		logger.Info("Application stopped graceful")
	}
}
