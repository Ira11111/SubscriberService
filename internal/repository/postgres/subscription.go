package postgres

import (
	t "SubscriberService/api/generated"
	"context"
)

func (s *Storage) SaveSub(ctx context.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil
}
func (s *Storage) GetSubs(ctx context.Context, limit int64, offset int64, subName string) ([]*t.Subscription, error) {
	return nil, nil
}
func (s *Storage) GetSubById(ctx context.Context, subId int64) (*t.Subscription, error) {
	return nil, nil
}
func (s *Storage) UpdateSub(ctx context.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil
}
func (s *Storage) DeleteSub(ctx context.Context, subId int64) error { return nil }
