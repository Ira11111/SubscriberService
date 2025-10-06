package service

import (
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
	SaveUserSub(ctx context.Context, userSub *d.SubscriptionUserCreate) (*d.SubscriptionUser, error)
	GetUserSubs(
		ctx context.Context,
		limit int64,
		offset int64,
		subName string,
		startDate time.Time,
		endDate time.Time,
	) ([]*d.SubscriptionUser, error)
	GetUsersForSub(
		ctx context.Context,
		subId int64,
		limit int64,
		offset int64,
		startDate time.Time,
		endDate time.Time) ([]*d.SubscriptionUser, error)
	GetSubsForUser(
		ctx context.Context,
		userid string,
		limit int64,
		offset int64,
		subName string,
		startDate time.Time,
		endDate time.Time,
	) ([]*d.SubscriptionUser, error)
}

type SubscriptionIdUserIdProvider interface {
	GetUserSubById(ctx context.Context, userId string, subId int64) (*d.SubscriptionUser, error)
	UpdateUserSub(ctx context.Context, userId string, subId int64, userSub *d.SubscriptionUserCreate) (*d.SubscriptionUser, error)
	DeleteUserSub(ctx context.Context, userId string, subId int64) error
	GetUserTotal(ctx context.Context, userId string, startDate time.Time, endDate time.Time) (*d.SubSum, error)
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
