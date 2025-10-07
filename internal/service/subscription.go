package service

import (
	t "SubscriberService/api/generated"
	"SubscriberService/internal/converter"
	d "SubscriberService/internal/domains"
	"context"
)

func (s *SubService) SaveSub(ctx context.Context, sub *t.Subscription) (*t.Subscription, error) {
	const op = "service.SaveSub"
	logger := s.logger.With("op", op)

	logger.Debug("Converting api type into domain type")
	domainSub := converter.ToDomainSubscription(sub)

	logger.Debug("trying to save domain type into DB")
	res, err := s.subProvider.SaveSub(ctx, domainSub)
	if err != nil {
		logger.Error("Failed to save subscription")
		return nil, err
	}
	logger.Debug("Convert domain type into api type")
	newSub := converter.ToAPISubscription(res)

	logger.Info("subscription saved successful")
	return newSub, nil
}
func (s *SubService) GetSubs(ctx context.Context, limit *t.LimitParam, offset *t.OffsetParam, subName *t.SubNameParam) ([]*t.Subscription, error) {
	const op = "service.GetSubs"
	logger := s.logger.With("op", op)
	var subs []d.Subscription
	var err error

	logger.Debug("parsing pagination params")
	lim, off := parsePagination(limit, offset)
	if subName != nil {
		logger.Debug("SubName is not empty, use GetSubsName")
		subs, err = s.subProvider.GetSubsName(ctx, lim, off, *subName)
	} else {
		logger.Debug("SubName is empty, use GetSubs")
		subs, err = s.subProvider.GetSubs(ctx, lim, off)
	}

	if err != nil {
		logger.Error("Failed to get subscriptions")
		return nil, err
	}
	logger.Debug("Converting models")
	apiSubs := converter.ToAPISubscriptionSlice(subs)

	logger.Info("Get subs successful")
	return apiSubs, nil
}
func (s *SubService) GetSubById(ctx context.Context, subId t.IdSubParam) (*t.Subscription, error) {
	const op = "service.subscription.GetSubById"
	logger := s.logger.With("op", op)

	logger.Debug("Trying to find sub")
	sub, err := s.subProvider.GetSubById(ctx, subId)
	if err != nil {
		logger.Error("Failed to find sub")
		return nil, err
	}
	logger.Debug("Converting domain type into api type")
	apiSub := converter.ToAPISubscription(sub)
	logger.Info("Find sub successful")
	return apiSub, nil

}
func (s *SubService) UpdateSub(ctx context.Context, sub *t.Subscription, subId t.IdSubParam) (*t.Subscription, error) {
	const op = "service.subscription.UpdateSub"
	logger := s.logger.With("op", op)

	// не меняем id и избегаем нулевого значения
	sub.SubId = &subId

	logger.Debug("converting api type into domain type")
	domSub := converter.ToDomainSubscription(sub)

	logger.Debug("Trying to update sub")
	updatedSub, err := s.subProvider.UpdateSub(ctx, domSub)
	if err != nil {
		logger.Error("Failed to update sub")
		return nil, err
	}
	logger.Debug("Converting domain type into api type")
	apiSub := converter.ToAPISubscription(updatedSub)
	logger.Info("Update sub successful")
	return apiSub, nil

}
func (s *SubService) DeleteSub(ctx context.Context, subId t.IdSubParam) error {
	const op = "service.subscription.DeleteSub"
	logger := s.logger.With("op", op)

	logger.Debug("Trying to delete sub")
	err := s.subProvider.DeleteSub(ctx, subId)
	if err != nil {
		logger.Error("Failed to delete sub")
		return err
	}
	logger.Info("Sub deleted")
	return nil
}
