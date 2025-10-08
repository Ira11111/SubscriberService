package service

import (
	t "SubscriberService/api/generated"
	d "SubscriberService/internal/domains"
	"SubscriberService/internal/filter"
	"context"
	"errors"
	"log/slog"
	"time"
)

const (
	limitDefault  = 20
	offsetDefault = 0
)

var (
	ErrNotfound              = errors.New("Not found")
	ErrOperationFailed       = errors.New("Operation failed")
	ErrFailedGetResponseData = errors.New("Failed to get response data")
)

type SubscriptionProvider interface {
	SaveSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error)
	GetSubs(ctx context.Context, options *filter.FilterOptions) ([]d.Subscription, error)
	GetSubById(ctx context.Context, subId int64) (*d.Subscription, error)
	UpdateSub(ctx context.Context, sub *d.Subscription) (*d.Subscription, error)
	DeleteSub(ctx context.Context, subId int64) error
}

type SubscriptionUserProvider interface {
	SaveUserSub(ctx context.Context, userSub *d.SubscriptionUser) (*d.SubscriptionUser, error)
	GetUserSubs(ctx context.Context, options *filter.FilterOptions) ([]d.SubscriptionUser, error)
}

type SubscriptionIdUserIdProvider interface {
	GetUserSubById(ctx context.Context, options *filter.FilterOptions) (*d.SubscriptionUser, error)
	UpdateUserSub(ctx context.Context, options *filter.FilterOptions, userSub *d.SubscriptionUser) (*d.SubscriptionUser, error)
	DeleteUserSub(ctx context.Context, options *filter.FilterOptions) error
	GetUserTotal(ctx context.Context, options *filter.FilterOptions) (int64, error)
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

func parsePagination(limit *int64, offset *int64) (int64, int64) {
	var l, off int64

	if limit != nil {
		l = *limit
	} else {
		l = limitDefault
	}

	if offset != nil {
		off = *offset
	} else {
		off = offsetDefault
	}
	return l, off
}

func parseDate(startDate *t.StartDateParam, endDate *t.EndDateParam) (time.Time, time.Time) {
	var startTime, endTime time.Time

	if startDate != nil {
		startTime = startDate.Time
	}

	if endDate != nil {
		endTime = endDate.Time
	}
	return startTime, endTime
}
