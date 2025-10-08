package postgres

import (
	d "SubscriberService/internal/domains"
	"SubscriberService/internal/filter"
	"SubscriberService/internal/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (s *Storage) GetUserSubById(ctx context.Context, options *filter.FilterOptions) (*d.SubscriptionUser, error) {
	query := `SELECT su.id_sub, su.id_user, su.start_date, su.end_date, s.name, s.price
	FROM subscription_user su INNER JOIN subscription s on s.id = su.id_sub`

	filteredQuery, args := filter.BuildQuery(query, options)

	var res d.SubscriptionUser

	err := sqlx.GetContext(ctx, s.db, &res, filteredQuery, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrDataNotFoud
		}
		return nil, repository.ErrFailedGet
	}
	return &res, nil

}
func (s *Storage) UpdateUserSub(ctx context.Context, options *filter.FilterOptions, userSub *d.SubscriptionUser) (*d.SubscriptionUser, error) {
	query := `UPDATE subscription_user su SET start_date = :start_date, end_date = :end_date 
            WHERE su.id_sub = :id_sub AND su.id_user = :id_user`

	rows, err := sqlx.NamedQueryContext(ctx, s.db, query, userSub)
	if err != nil {
		fmt.Println(err.Error())
		return nil, repository.ErrUpdateFailed
	}
	if !rows.Next() {
		return nil, repository.ErrDataNotFoud
	}

	// находим и возвращаем обновленную запись
	return s.GetUserSubById(ctx, options)

}
func (s *Storage) DeleteUserSub(ctx context.Context, options *filter.FilterOptions) error {
	query := `DELETE FROM subscription_user su`
	filteredQuery, args := filter.BuildQuery(query, options)

	result, err := s.db.ExecContext(ctx, filteredQuery, args...)

	if err != nil {
		return repository.ErrFailedDelete
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return repository.ErrDataNotFoud
	}

	return nil
}
func (s *Storage) GetUserTotal(ctx context.Context, options *filter.FilterOptions) (int64, error) {
	query := `SELECT COALESCE(SUM(s.price), 0) FROM subscription s INNER JOIN subscription_user su ON s.id = su.id_sub`
	filteredQuery, args := filter.BuildQuery(query, options)

	var sum int64

	err := sqlx.GetContext(ctx, s.db, &sum, filteredQuery, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, repository.ErrDataNotFoud
		}
		return 0, repository.ErrFailedGet
	}
	return sum, nil
}
