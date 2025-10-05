package postgres

import (
	"SubscriberService/internal/config"
	"SubscriberService/internal/repository"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(cfg *config.DBConfig) (*Storage, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s host=%s dbname=%s password=%s sslmode=%s",
			cfg.User, cfg.Host, cfg.Database, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, repository.ErrFailedConnect
	}
	if err = db.Ping(); err != nil {
		return nil, repository.ErrFailedConnect
	}
	return &Storage{db: db}, nil
}
