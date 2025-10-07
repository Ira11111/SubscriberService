package service

import (
	t "SubscriberService/api/generated"
	"context"
)

func (s *SubService) GetUserSubById(ctx context.Context, userId t.IdUserParam, subId t.IdSubParam) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *SubService) UpdateUserSub(ctx context.Context, userId t.IdUserParam, subId t.IdSubParam, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *SubService) DeleteUserSub(ctx context.Context, userId t.IdUserParam, subId t.IdSubParam) error {
	return nil
}
func (s *SubService) GetUserTotal(ctx context.Context, userId t.IdUserParam, startDate *t.StartDateParam, endDate *t.EndDateParam) (*t.SubSum, error) {
	return nil, nil
}
