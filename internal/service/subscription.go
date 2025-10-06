package service

import (
	t "SubscriberService/api/generated"
	"SubscriberService/internal/converter"
	d "SubscriberService/internal/domains"
	"context"
)

func (s *SubService) SaveSub(ctx context.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil
}
func (s *SubService) GetSubs(ctx context.Context, limit t.LimitParam, offset t.OffsetParam, subName t.SubNameParam) ([]*t.Subscription, error) {
	var subs []*d.Subscription
	var err error
	if subName != "" {
		s.logger.Debug("SubName is not empty, use GetSubsName")
		subs, err = s.subProvider.GetSubsName(ctx, limit, offset, subName)
	} else {
		s.logger.Debug("SubName is empty, use GetSubs")
		subs, err = s.subProvider.GetSubs(ctx, limit, offset)
	}

	if err != nil {
		s.logger.Error("Failed to get subscriptions")
		return nil, err
	}
	s.logger.Info("Get subs successful")
	s.logger.Debug("Converting models")
	apiSubs := converter.ToAPISubscriptionSlice(subs)
	return apiSubs, nil
}
func (s *SubService) GetSubById(ctx context.Context, subId t.IdSubParam) (*t.Subscription, error) {
	return nil, nil

}
func (s *SubService) UpdateSub(ctx context.Context, sub *t.Subscription) (*t.Subscription, error) {
	return nil, nil

}
func (s *SubService) DeleteSub(ctx context.Context, subId t.IdSubParam) error {
	return nil

}
