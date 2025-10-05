package service

import (
	t "SubscriberService/api/generated"

	"github.com/go-chi/chi/v5"
)

func (s *SubService) GetUserSubById(ctx *chi.Context, userId *t.IdUserParam, subId *t.IdSubParam) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *SubService) UpdateUserSub(ctx *chi.Context, userId *t.IdUserParam, subId *t.IdSubParam, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *SubService) DeleteUserSub(ctx *chi.Context, userId *t.IdUserParam, subId *t.IdSubParam) error {
	return nil
}
func (s *SubService) GetUserTotal(ctx *chi.Context, userId *t.IdUserParam, startDate *t.StartDateParam, endDate *t.EndDateParam) (*t.SubSum, error) {
	return nil, nil
}
