package postgres

import (
	d "SubscriberService/internal/domains"
	"context"
)

func (s *Storage) SaveSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error) {
	return nil, nil
}
func (s *Storage) GetSubs(ctx context.Context, limit int64, offset int64) ([]*d.Subscription, error) {
	return nil, nil
}
func (s *Storage) GetSubsName(ctx context.Context, limit int64, offset int64, subName string) ([]*d.Subscription, error) {
	return nil, nil
}

func (s *Storage) GetSubById(ctx context.Context, subId int64) (*d.Subscription, error) {
	return nil, nil
}
func (s *Storage) UpdateSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error) {
	return nil, nil
}
func (s *Storage) DeleteSub(ctx context.Context, subId int64) error { return nil }
