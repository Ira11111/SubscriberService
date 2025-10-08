package postgres

import (
	d "SubscriberService/internal/domains"
	"SubscriberService/internal/filter"
	"SubscriberService/internal/repository"
	"context"
	"database/sql"
	"errors"

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
		return nil, repository.ErrFailedSave
	}
	defer rows.Close()

	var newSub d.Subscription
	if rows.Next() {
		if err = rows.Scan(&newSub.Id, &newSub.ServiceName, &newSub.Price); err != nil {
			return nil, repository.ErrFailedScan
		}
	}
	return &newSub, nil
}
func (s *Storage) GetSubs(ctx context.Context, options *filter.FilterOptions) ([]d.Subscription, error) {
	var subs []d.Subscription

	query := `SELECT * FROM subscription`

	filteredQuery, args := filter.BuildQuery(query, options)
	err := sqlx.SelectContext(ctx, s.db, &subs, filteredQuery, args...)
	if err != nil {
		return nil, repository.ErrFailedGet
	}
	return subs, nil
}

func (s *Storage) GetSubById(ctx context.Context, subId int64) (*d.Subscription, error) {
	var sub d.Subscription
	query := `SELECT * FROM subscription WHERE id = $1`
	err := sqlx.GetContext(ctx, s.db, &sub, query, subId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrDataNotFoud
		}
		return nil, repository.ErrFailedGet
	}
	return &sub, nil
}

func (s *Storage) UpdateSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error) {
	var updatedSub d.Subscription
	query := `UPDATE subscription SET id = :id, name = :name, price = :price 
                    WHERE id = :id RETURNING id, name, price`

	rows, err := sqlx.NamedQueryContext(ctx, s.db, query, sub)
	if err != nil {
		return nil, repository.ErrUpdateFailed
	}
	if !rows.Next() {
		return nil, repository.ErrDataNotFoud
	}
	if err = rows.Scan(&updatedSub.Id, &updatedSub.ServiceName, &updatedSub.Price); err != nil {
		return nil, err
	}
	return &updatedSub, nil
}

func (s *Storage) DeleteSub(ctx context.Context, subId int64) error {
	query := `DELETE FROM subscription WHERE id = $1`

	result, err := s.db.ExecContext(ctx, query, subId)
	if err != nil {
		return repository.ErrFailedDelete
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return repository.ErrDataNotFoud
	}

	return nil
}
