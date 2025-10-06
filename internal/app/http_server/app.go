package http_server

import (
	"SubscriberService/api/generated"
	"SubscriberService/internal/config"
	"SubscriberService/internal/http/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ServerApp struct {
	Server *http.Server
}

func NewServer(cfg *config.ServerConfig, handler *handler.Handler) *ServerApp {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	if cfg.Middleware.Logger.Enabled {
		r.Use(middleware.Logger)
	}
	if cfg.Middleware.Recovery.Enabled {
		r.Use(middleware.Recoverer)
	}
	if cfg.Middleware.Timeout.Enabled {
		r.Use(middleware.Timeout(cfg.Middleware.Timeout.Duration))
	}

	//r.Get("/health", healthCheckHandler)

	// Подключаем сгенерированные роуты
	generated.HandlerFromMux(handler, r)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &ServerApp{
		Server: srv,
	}
}
