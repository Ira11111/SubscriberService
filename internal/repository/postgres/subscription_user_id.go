package postgres

import (
	t "SubscriberService/api/generated"
	"time"

	"github.com/go-chi/chi/v5"
)

func (s *Storage) GetUserSubById(ctx *chi.Context, userId string, subId int64) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) UpdateUserSub(ctx *chi.Context, userId string, subId int64, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) DeleteUserSub(ctx *chi.Context, userId string, subId int64) error { return nil }
func (s *Storage) GetUserTotal(ctx *chi.Context, userId string, startDate *time.Time, endDate time.Time) (*t.SubSum, error) {
	return nil, nil
}
