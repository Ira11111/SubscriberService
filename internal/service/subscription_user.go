package service

import (
	t "SubscriberService/api/generated"
	"SubscriberService/internal/converter"
	"SubscriberService/internal/filter"
	"SubscriberService/internal/repository"
	"context"
	"errors"
)

func (s *SubService) SaveUserSub(ctx context.Context, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error) {
	const op = "service.subscription_user.SaveUserSub"
	logger := s.logger.With("op", op)

	logger.Debug("Converting into domain type")
	domainUserSub := converter.ToDomainSubscriptionUserCreate(userSub)

	logger.Debug("trying to save domain type into DB")
	res, err := s.subUserProvider.SaveUserSub(ctx, domainUserSub)

	if err != nil {
		if errors.Is(err, repository.ErrFailedSave) {
			logger.Error("Failed to save user subscription")
			return nil, ErrOperationFailed
		}
		if errors.Is(err, repository.ErrFailedScan) {
			logger.Warn("Failed to get response data")
			return nil, ErrFailedGetResponseData
		}
		return nil, err
	}
	logger.Debug("Convert domain type into api type")
	newSub := converter.ToAPISubscriptionUser(res)

	logger.Info("subscription saved successful")
	return newSub, nil
}

func (s *SubService) GetUserSubs(
	ctx context.Context,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	subName *t.SubNameParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam,
) ([]*t.SubscriptionUser, error) {
	const op = "service.subscription_user.GetUserSubs"
	logger := s.logger.With("op", op)

	logger.Debug("parsing pagination params")
	lim, off := parsePagination(limit, offset)

	f := filter.NewFilterBuilder().
		WithPagination(lim, off).
		WithSubName(subName).
		WithDateRange(startDate.Time, endDate.Time).
		Build()

	userSubs, err := s.subUserProvider.GetUserSubs(ctx, &f)

	if err != nil {
		logger.Error("Failed to get user subscriptions")
		return nil, ErrOperationFailed
	}
	if len(userSubs) == 0 {
		return nil, ErrNotfound
	}

	logger.Debug("Converting models")
	apiSubs := converter.ToAPISubscriptionUserSlice(userSubs)

	logger.Info("Get subs successful")
	return apiSubs, nil

}
func (s *SubService) GetUsersForSub(
	ctx context.Context,
	subId t.IdSubParam,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam) ([]*t.SubscriptionUser, error) {

	const op = "service.subscription_user.GetUsersForSub"
	logger := s.logger.With("op", op)

	lim, off := parsePagination(limit, offset)
	f := filter.NewFilterBuilder().
		WithPagination(lim, off).
		WithSubID(subId).
		WithDateRange(startDate.Time, endDate.Time).
		Build()

	userSubs, err := s.subUserProvider.GetUserSubs(ctx, &f)

	if err != nil {
		logger.Error("Failed to get user subscriptions")
		return nil, ErrOperationFailed
	}
	if len(userSubs) == 0 {
		return nil, ErrNotfound
	}

	logger.Debug("Converting models")
	apiSubs := converter.ToAPISubscriptionUserSlice(userSubs)

	logger.Info("Get subs successful")
	return apiSubs, nil

}
func (s *SubService) GetSubsForUser(
	ctx context.Context,
	userId t.IdUserParam,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	subName *t.SubNameParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam,
) ([]*t.SubscriptionUser, error) {
	const op = "service.subscription_user.GetSubsForUser"
	logger := s.logger.With("op", op)

	lim, off := parsePagination(limit, offset)
	f := filter.NewFilterBuilder().
		WithPagination(lim, off).
		WithUserID(userId).
		WithSubName(subName).
		WithDateRange(startDate.Time, endDate.Time).
		Build()

	userSubs, err := s.subUserProvider.GetUserSubs(ctx, &f)

	if err != nil {
		logger.Error("Failed to get user subscriptions")
		return nil, ErrOperationFailed
	}
	if len(userSubs) == 0 {
		return nil, ErrNotfound
	}

	logger.Debug("Converting models")
	apiSubs := converter.ToAPISubscriptionUserSlice(userSubs)

	logger.Info("Get subs successful")
	return apiSubs, nil
}
