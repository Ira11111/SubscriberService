package service

import (
	t "SubscriberService/api/generated"
	d "SubscriberService/internal/domains"
	"context"
	"log/slog"
	"time"
)

type SubscriptionProvider interface {
	SaveSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error)
	GetSubs(ctx context.Context, limit int64, offset int64) ([]*d.Subscription, error)
	GetSubsName(ctx context.Context, limit int64, offset int64, subName string) ([]*d.Subscription, error)
	GetSubById(ctx context.Context, subId int64) (*d.Subscription, error)
	UpdateSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error)
	DeleteSub(ctx context.Context, subId int64) error
}

type SubscriptionUserProvider interface {
	SaveUserSub(ctx context.Context, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error)
	GetUserSubs(
		ctx context.Context,
		limit int64,
		offset int64,
		subName string,
		startDate time.Time,
		endDate time.Time,
	) ([]*t.SubscriptionUser, error)
	GetUsersForSub(
		ctx context.Context,
		subId int64,
		limit int64,
		offset int64,
		startDate time.Time,
		endDate time.Time) ([]*t.SubscriptionUser, error)
	GetSubsForUser(
		ctx context.Context,
		userid string,
		limit int64,
		offset int64,
		subName string,
		startDate time.Time,
		endDate time.Time,
	) ([]*t.SubscriptionUser, error)
}

type SubscriptionIdUserIdProvider interface {
	GetUserSubById(ctx context.Context, userId string, subId int64) (*t.SubscriptionUser, error)
	UpdateUserSub(ctx context.Context, userId string, subId int64, userSub *t.SubscriptionUserCreate) (*t.SubscriptionUser, error)
	DeleteUserSub(ctx context.Context, userId string, subId int64) error
	GetUserTotal(ctx context.Context, userId string, startDate *time.Time, endDate time.Time) (*t.SubSum, error)
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
