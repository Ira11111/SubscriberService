package service

import (
	t "SubscriberService/api/generated"

	"github.com/go-chi/chi/v5"
)

func (s *SubService) SaveSub(ctx *chi.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil
}
func (s *SubService) GetSubs(ctx *chi.Context, limit *t.LimitParam, offset *t.OffsetParam, subName *t.SubNameParam) ([]*t.Subscription, error) {
	return nil, nil
}
func (s *SubService) GetSubById(ctx *chi.Context, subId *t.IdSubParam) (*t.Subscription, error) {
	return nil, nil

}
func (s *SubService) UpdateSub(ctx *chi.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil

}
func (s *SubService) DeleteSub(ctx *chi.Context, subId *t.IdSubParam) error {
	return nil

}
