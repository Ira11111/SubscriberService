package postgres

import (
	d "SubscriberService/internal/domains"
	"context"
	"time"
)

func (s *Storage) GetUserSubById(ctx context.Context, userId string, subId int64) (*d.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) UpdateUserSub(ctx context.Context, userId string, subId int64, userSub *d.SubscriptionUserCreate) (*d.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) DeleteUserSub(ctx context.Context, userId string, subId int64) error { return nil }
func (s *Storage) GetUserTotal(ctx context.Context, userId string, startDate time.Time, endDate time.Time) (*d.SubSum, error) {
	return nil, nil
}
