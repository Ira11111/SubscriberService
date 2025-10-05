package service

import (
	t "SubscriberService/api/generated"

	"github.com/go-chi/chi/v5"
)

func (s *SubService) SaveUserSub(ctx *chi.Context, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error) {
	return nil, nil

}
func (s *SubService) GetUserSubs(
	ctx *chi.Context,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	subName *t.SubNameParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam,
) ([]*t.SubscriptionUser, error) {
	return nil, nil

}
func (s *SubService) GetUsersForSub(
	ctx *chi.Context,
	subId *t.IdSubParam,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam) ([]*t.SubscriptionUser, error) {
	return nil, nil

}
func (s *SubService) GetSubsForUser(
	ctx *chi.Context,
	userid *t.IdUserParam,
	limit *t.LimitParam,
	offset *t.OffsetParam,
	subName *t.SubNameParam,
	startDate *t.StartDateParam,
	endDate *t.EndDateParam,
) ([]*t.SubscriptionUser, error) {
	return nil, nil
}
