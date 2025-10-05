package postgres

import (
	t "SubscriberService/api/generated"

	"github.com/go-chi/chi/v5"
)

func (s *Storage) SaveSub(ctx *chi.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil
}
func (s *Storage) GetSubs(ctx *chi.Context, limit int64, offset int64, subName string) ([]*t.Subscription, error) {
	return nil, nil
}
func (s *Storage) GetSubById(ctx *chi.Context, subId int64) (*t.Subscription, error) { return nil, nil }
func (s *Storage) UpdateSub(ctx *chi.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil
}
func (s *Storage) DeleteSub(ctx *chi.Context, subId int64) error { return nil }
