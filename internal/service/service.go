package service

import (
	t "SubscriberService/api/generated"
	"log/slog"
	"time"

	"github.com/go-chi/chi/v5"
)

type SubscriptionProvider interface {
	SaveSub(ctx *chi.Context, sub *t.Subscription) (*t.Subscription, error)
	GetSubs(ctx *chi.Context, limit int64, offset int64, subName string) ([]*t.Subscription, error)
	GetSubById(ctx *chi.Context, subId int64) (*t.Subscription, error)
	UpdateSub(ctx *chi.Context, sub *t.Subscription) (*t.Subscription, error)
	DeleteSub(ctx *chi.Context, subId int64) error
}

type SubscriptionUserProvider interface {
	SaveUserSub(ctx *chi.Context, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error)
	GetUserSubs(
		ctx *chi.Context,
		limit int64,
		offset int64,
		subName string,
		startDate time.Time,
		endDate time.Time,
	) ([]*t.SubscriptionUser, error)
	GetUsersForSub(
		ctx *chi.Context,
		subId int64,
		limit int64,
		offset int64,
		startDate time.Time,
		endDate time.Time) ([]*t.SubscriptionUser, error)
	GetSubsForUser(
		ctx *chi.Context,
		userid string,
		limit int64,
		offset int64,
		subName string,
		startDate time.Time,
		endDate time.Time,
	) ([]*t.SubscriptionUser, error)
}

type SubscriptionIdUserIdProvider interface {
	GetUserSubById(ctx *chi.Context, userId string, subId int64) (*t.SubscriptionUser, error)
	UpdateUserSub(ctx *chi.Context, userId string, subId int64, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error)
	DeleteUserSub(ctx *chi.Context, userId string, subId int64) error
	GetUserTotal(ctx *chi.Context, userId string, startDate *time.Time, endDate time.Time) (*t.SubSum, error)
}

type SubService struct {
	logger              *slog.Logger
	subProvider         SubscriptionProvider
	subUserProvider     SubscriptionUserProvider
	subIdUserIdProvider SubscriptionIdUserIdProvider
}

func NewSubService(
	logger *slog.Logger,
	subProvider SubscriptionProvider,
	subUserProvider SubscriptionUserProvider,
	subIdUserIdProvider SubscriptionIdUserIdProvider) *SubService {

	return &SubService{
		logger:              logger,
		subProvider:         subProvider,
		subUserProvider:     subUserProvider,
		subIdUserIdProvider: subIdUserIdProvider,
	}
}
