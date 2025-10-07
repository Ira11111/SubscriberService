package postgres

import (
	d "SubscriberService/internal/domains"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (s *Storage) SaveSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error) {
	query := `
        INSERT INTO subscription (name, price) 
        VALUES (:name, :price) 
        RETURNING id, name, price
    `

	rows, err := sqlx.NamedQueryContext(ctx, s.db, query, sub)
	if err != nil {
		return nil, fmt.Errorf("create subscription: %w", err)
	}
	defer rows.Close()

	var newSub d.Subscription
	if rows.Next() {
		if err = rows.Scan(&newSub.Id, &newSub.ServiceName, &newSub.Price); err != nil {
			return nil, fmt.Errorf("scan created subscription: %w", err)
		}
	}

	return &newSub, nil
}
func (s *Storage) GetSubs(ctx context.Context, limit int64, offset int64) ([]d.Subscription, error) {
	var subs []d.Subscription

	query := `SELECT * FROM subscription LIMIT $1 OFFSET $2`
	err := sqlx.SelectContext(ctx, s.db, &subs, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (s *Storage) GetSubsName(ctx context.Context, limit int64, offset int64, subName string) ([]d.Subscription, error) {
	var subs []d.Subscription
	searchPattern := "%" + subName + "%"
	query := `SELECT * FROM subscription s WHERE s.name ILIKE $1 LIMIT $2 OFFSET $3`
	err := sqlx.SelectContext(ctx, s.db, &subs, query, searchPattern, limit, offset)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (s *Storage) GetSubById(ctx context.Context, subId int64) (*d.Subscription, error) {
	return nil, nil
}
func (s *Storage) UpdateSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error) {
	return nil, nil
}
func (s *Storage) DeleteSub(ctx context.Context, subId int64) error { return nil }
