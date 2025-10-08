package postgres

import (
	d "SubscriberService/internal/domains"
	"SubscriberService/internal/filter"
	"SubscriberService/internal/repository"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (s *Storage) SaveUserSub(ctx context.Context, userSub *d.SubscriptionUser) (*d.SubscriptionUser, error) {
	query := `
        INSERT INTO subscription_user (id_sub, id_user, start_date) 
        VALUES (:id_sub, :id_user, :start_date) 
        RETURNING id_sub, id_user, start_date
    `

	rows, err := sqlx.NamedQueryContext(ctx, s.db, query, userSub)
	if err != nil {
		// проверка foreign key
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23503" {
				return nil, repository.ErrDataNotFoud
			}
		}
		return nil, repository.ErrFailedSave
	}
	defer rows.Close()

	var newSub d.SubscriptionUser
	if rows.Next() {
		if err = rows.Scan(&newSub.SubId, &newSub.UserId, &newSub.StartDate); err != nil {
			return nil, repository.ErrFailedScan
		}
	}
	return &newSub, nil
}

func (s *Storage) GetUserSubs(ctx context.Context, options *filter.FilterOptions) ([]d.SubscriptionUser, error) {
	query := `SELECT su.id_sub, su.id_user, su.start_date, su.end_date, s.name, s.price
	FROM subscription_user su INNER JOIN subscription s on s.id = su.id_sub`
	filteredQuery, args := filter.BuildQuery(query, options)
	var res []d.SubscriptionUser

	err := sqlx.SelectContext(ctx, s.db, &res, filteredQuery, args...)
	if err != nil {
		return nil, repository.ErrFailedGet
	}
	return res, nil
}
