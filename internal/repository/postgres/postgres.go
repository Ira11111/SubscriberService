package postgres

import (
	"SubscriberService/internal/config"
	"SubscriberService/internal/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const maxRetries = 5

type Storage struct {
	db *sqlx.DB
}

func NewStorage(cfg *config.DBConfig) (*Storage, error) {
	dsn := fmt.Sprintf("user=%s host=%s dbname=%s password=%s sslmode=%s",
		cfg.User, cfg.Host, cfg.Database, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open("postgres", dsn)
	for i := 0; i < maxRetries; i++ {
		if err != nil {
			time.Sleep(2 * time.Second)
			db, err = sqlx.Open("postgres", dsn)
		} else {
			break
		}
	}

	if err != nil {
		return nil, repository.ErrFailedConnect
	}
	if err = db.Ping(); err != nil {
		return nil, repository.ErrFailedConnect
	}
	return &Storage{db: db}, nil
}

func (s *Storage) CLose(ctx context.Context) error {
	var err error

	done := make(chan struct{})

	go func() {
		defer close(done)
		// Запрещаем новые соединения
		s.db.SetMaxOpenConns(0)
		s.db.SetMaxIdleConns(0)

		// Ждем завершения активных операций (если нужно)
		stats := s.db.Stats()
		if stats.InUse > 0 {
			time.Sleep(2 * time.Second)
		}
		// Закрываем соединение
		err = s.db.Close()
	}()

	select {
	case <-done:
		return err
	case <-ctx.Done():
		// при отмене контекста принудительно зарываем БД
		s.db.Close()
		return errors.New("Failed to stop db graceful")
	}
}
