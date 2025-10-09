package service

import (
	t "SubscriberService/api/generated"
	"SubscriberService/internal/converter"
	"SubscriberService/internal/filter"
	"SubscriberService/internal/repository"
	"context"
	"errors"
	"fmt"
)

func (s *SubService) GetUserSubById(ctx context.Context, userId *t.IdUserParam, subId t.IdSubParam) (*t.SubscriptionUser, error) {
	const op = "service.subscription_user_id.GetUserSubById"
	logger := s.logger.With("op", op)

	logger.Debug("Trying to find user sub")
	f := filter.NewFilterBuilder().
		WithEqualCondition("su.id_user", userId).
		WithEqualCondition("su.id_sub", subId).
		Build()
	data, err := s.subIdUserIdProvider.GetUserSubById(ctx, &f)

	if err != nil {
		logger.Error("Failed to find sub")
		if errors.Is(err, repository.ErrDataNotFoud) {
			return nil, ErrNotfound
		}
		return nil, ErrOperationFailed
	}

	logger.Debug("Converting domain type into api type")
	apiSub := converter.ToAPISubscriptionUser(data)
	logger.Info("Find sub successful")
	return apiSub, nil
}
func (s *SubService) UpdateUserSub(ctx context.Context, userId t.IdUserParam, subId t.IdSubParam, userSub *t.SubscriptionUserUpdate) (*t.SubscriptionUser, error) {
	const op = "service.subscription_user_id.UpdateUserSub"
	logger := s.logger.With("op", op)

	logger.Debug("converting api type into domain type")
	domSub := converter.ToDomainSubscriptionUserUpdate(userSub)
	domSub.UserId = userId
	domSub.SubId = subId

	logger.Debug("Trying to update sub")
	f := filter.NewFilterBuilder().
		WithEqualCondition("su.id_user", userId).
		WithEqualCondition("su.id_sub", subId).
		Build()
	updatedSub, err := s.subIdUserIdProvider.UpdateUserSub(ctx, &f, domSub)

	if err != nil {
		fmt.Println(err.Error())
		logger.Error("Failed to update sub")
		if errors.Is(err, repository.ErrUpdateFailed) {
			return nil, ErrOperationFailed
		}
		if errors.Is(err, repository.ErrDataNotFoud) {
			return nil, ErrNotfound
		}
		return nil, err
	}

	logger.Debug("Converting domain type into api type")
	apiData := converter.ToAPISubscriptionUser(updatedSub)
	logger.Info("Update successful")
	return apiData, nil
}
func (s *SubService) DeleteUserSub(ctx context.Context, userId *t.IdUserParam, subId t.IdSubParam) error {
	const op = "service.subscription_user_id.DeleteUserSub"
	logger := s.logger.With("op", op)

	logger.Debug("Trying to delete user sub")
	f := filter.NewFilterBuilder().
		WithEqualCondition("su.id_user", userId).
		WithEqualCondition("su.id_sub", subId).
		Build()
	err := s.subIdUserIdProvider.DeleteUserSub(ctx, &f)
	if err != nil {
		logger.Error("Failed to delete user sub")
		if errors.Is(err, repository.ErrDataNotFoud) {
			return ErrNotfound
		}
		if errors.Is(err, repository.ErrFailedDelete) {
			return ErrOperationFailed
		}
		return err
	}
	logger.Info("user subscription deleted")
	return nil
}
func (s *SubService) GetUserTotal(ctx context.Context, userId *t.IdUserParam, startDate *t.StartDateParam, endDate *t.EndDateParam) (*t.SubSum, error) {
	const op = "service.subscription_user_id.GetUserTotal"
	logger := s.logger.With("op", op)

	logger.Debug("Trying to find user sub")
	st, end := parseDate(startDate, endDate)
	f := filter.NewFilterBuilder().
		WithEqualCondition("su.id_user", userId).
		WithDateRange(st, end).Build()
	res, err := s.subIdUserIdProvider.GetUserTotal(ctx, &f)

	if err != nil {
		logger.Error("Failed to find sub")
		if errors.Is(err, repository.ErrDataNotFoud) {
			return nil, ErrNotfound
		}
		return nil, ErrOperationFailed
	}

	logger.Debug("Converting domain type into api type")
	apiData := &t.SubSum{
		UserId:    userId,
		StartDate: startDate,
		EndDate:   endDate,
		TotalSum:  &res,
	}
	logger.Info("Find sub successful")
	return apiData, nil
}
