package service

import (
	t "SubscriberService/api/generated"
	"context"
)

func (s *SubService) SaveSub(ctx context.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil
}
func (s *SubService) GetSubs(ctx context.Context, limit *t.LimitParam, offset *t.OffsetParam, subName *t.SubNameParam) ([]*t.Subscription, error) {
	return nil, nil
}
func (s *SubService) GetSubById(ctx context.Context, subId *t.IdSubParam) (*t.Subscription, error) {
	return nil, nil

}
func (s *SubService) UpdateSub(ctx context.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil

}
func (s *SubService) DeleteSub(ctx context.Context, subId *t.IdSubParam) error {
	return nil

}
