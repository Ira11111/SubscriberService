package service

import (
	t "SubscriberService/api/generated"
	"context"
)

func (s *SubService) SaveUserSub(ctx context.Context, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error) {
	return nil, nil

}
func (s *SubService) GetUserSubs(
	ctx context.Context,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	subName *t.SubNameParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam,
) ([]*t.SubscriptionUser, error) {
	return nil, nil

}
func (s *SubService) GetUsersForSub(
	ctx context.Context,
	subId *t.IdSubParam,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam) ([]*t.SubscriptionUser, error) {
	return nil, nil

}
func (s *SubService) GetSubsForUser(
	ctx context.Context,
	userid *t.IdUserParam,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	subName *t.SubNameParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam,
) ([]*t.SubscriptionUser, error) {
	return nil, nil
}
