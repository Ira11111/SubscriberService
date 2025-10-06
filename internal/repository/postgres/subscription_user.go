package postgres

import (
	d "SubscriberService/internal/domains"
	"context"
	"time"
)

func (s *Storage) SaveUserSub(ctx context.Context, userSub *d.SubscriptionUserCreate) (*d.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) GetUserSubs(
	ctx context.Context,
	limit int64,
	offset int64,
	subName string,
	startDate time.Time,
	endDate time.Time,
) ([]*d.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) GetUsersForSub(
	ctx context.Context,
	subId int64,
	limit int64,
	offset int64,
	startDate time.Time,
	endDate time.Time) ([]*d.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) GetSubsForUser(
	ctx context.Context,
	userid string,
	limit int64,
	offset int64,
	subName string,
	startDate time.Time,
	endDate time.Time,
) ([]*d.SubscriptionUser, error) {
	return nil, nil
}
