package postgres

import (
	t "SubscriberService/api/generated"
	"context"
	"time"
)

func (s *Storage) GetUserSubById(ctx context.Context, userId string, subId int64) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) UpdateUserSub(ctx context.Context, userId string, subId int64, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) DeleteUserSub(ctx context.Context, userId string, subId int64) error { return nil }
func (s *Storage) GetUserTotal(ctx context.Context, userId string, startDate *time.Time, endDate time.Time) (*t.SubSum, error) {
	return nil, nil
}
