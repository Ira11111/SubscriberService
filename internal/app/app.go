package app

import (
	"SubscriberService/internal/app/http_server"
	"SubscriberService/internal/config"
	"SubscriberService/internal/http/handler"
	"SubscriberService/internal/repository/postgres"
	"SubscriberService/internal/service"
	"log/slog"
)

type App struct {
	Server  *http_server.ServerApp
	Service *service.SubService
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	storage, err := postgres.NewStorage(&cfg.DB)
	if err != nil {
		panic(err.Error())
	}
	newService := service.NewSubService(logger, storage, storage, storage)
	newHandler := handler.NewHandler(logger, newService, newService, newService)
	newServer := http_server.NewServer(&cfg.HttpServer, newHandler)

	return &App{
		Server:  newServer,
		Service: newService,
	}
}
