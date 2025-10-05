package postgres

import (
	t "SubscriberService/api/generated"
	"time"

	"github.com/go-chi/chi/v5"
)

func (s *Storage) SaveUserSub(ctx *chi.Context, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) GetUserSubs(
	ctx *chi.Context,
	limit int64,
	offset int64,
	subName string,
	startDate time.Time,
	endDate time.Time,
) ([]*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) GetUsersForSub(
	ctx *chi.Context,
	subId int64,
	limit int64,
	offset int64,
	startDate time.Time,
	endDate time.Time) ([]*t.SubscriptionUser, error) {
	return nil, nil
}
func (s *Storage) GetSubsForUser(
	ctx *chi.Context,
	userid string,
	limit int64,
	offset int64,
	subName string,
	startDate time.Time,
	endDate time.Time,
) ([]*t.SubscriptionUser, error) {
	return nil, nil
}
