package service

import (
	t "SubscriberService/api/generated"
	"SubscriberService/internal/converter"
	d "SubscriberService/internal/domains"
	"context"
)

func (s *SubService) SaveSub(ctx context.Context, sub *t.Subscription) (*t.Subscription, error) {
	domainSub := converter.ToDomainSubscription(sub)
	res, err := s.subProvider.SaveSub(ctx, domainSub)
	if err != nil {
		return nil, err
	}

	newSub := converter.ToAPISubscription(res)
	return newSub, nil
}
func (s *SubService) GetSubs(ctx context.Context, limit *t.LimitParam, offset *t.OffsetParam, subName *t.SubNameParam) ([]*t.Subscription, error) {
	var subs []d.Subscription
	var err error
	lim, off := parsePagination(limit, offset)
	if subName != nil {
		s.logger.Debug("SubName is not empty, use GetSubsName")
		subs, err = s.subProvider.GetSubsName(ctx, lim, off, *subName)
	} else {
		s.logger.Debug("SubName is empty, use GetSubs")
		subs, err = s.subProvider.GetSubs(ctx, lim, off)
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
