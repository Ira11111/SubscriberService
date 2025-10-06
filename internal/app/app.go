package app

import (
	"SubscriberService/internal/app/http_server"
	"SubscriberService/internal/config"
	"SubscriberService/internal/http/handler"
	"SubscriberService/internal/repository/postgres"
	"SubscriberService/internal/service"
	"context"
	"log/slog"
)

type App struct {
	Server  *http_server.ServerApp
	Service *service.SubService
	Storage *postgres.Storage
	logger  *slog.Logger
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	storage, err := postgres.NewStorage(&cfg.DB)
	if err != nil {
		panic(err.Error())
	}
	logger.Info("Storage is up")
	newService := service.NewSubService(logger, storage, storage, storage)
	newHandler := handler.NewHandler(logger, newService, newService, newService)
	newServer := http_server.NewServer(&cfg.HttpServer, newHandler)

	return &App{
		Server:  newServer,
		Service: newService,
		Storage: storage,
		logger:  logger,
	}
}

func (a *App) Start() {
	a.logger.Info("Server is running")
	a.Server.Server.ListenAndServe()

}

func (a *App) Shutdown(ctx context.Context) error {
	a.logger.Info("Shutting down application components")

	// Останавливаем HTTP сервер
	if err := a.Server.Server.Shutdown(ctx); err != nil {
		a.logger.Warn("HTTP server shutdown warning", slog.Any("error", err))
		return err
	}
	a.logger.Info("Server stopped graceful")

	// Закрываем соединение с БД
	if a.Storage != nil {
		if err := a.Storage.CLose(ctx); err != nil {
			a.logger.Warn("Database connection close warning", slog.Any("error", err))
			return err
		}
		a.logger.Info("Database connection closed graceful")
	}

	return nil
}
