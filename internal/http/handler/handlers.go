package handler

import (
	api "SubscriberService/api/generated"
	"context"
	"log/slog"
)

/*
Paths:
/subscriptions
/subscriptions/id
*/
type SubscriptionService interface {
	SaveSub(ctx context.Context, sub *api.Subscription) (*api.Subscription, error)
	GetSubs(ctx context.Context, limit *api.LimitParam, offset *api.OffsetParam, subName *api.SubNameParam) ([]*api.Subscription, error)
	GetSubById(ctx context.Context, subId api.IdSubParam) (*api.Subscription, error)
	UpdateSub(ctx context.Context, sub *api.Subscription, subId api.IdSubParam) (*api.Subscription, error)
	DeleteSub(ctx context.Context, subId api.IdSubParam) error
}

/*
Paths:
/subscriptions/users
/subscriptions/id/users
/subscriptions/users/id
*/
type SubscriptionUserService interface {
	SaveUserSub(ctx context.Context, userSub *api.SubscriptionUser) (*api.SubscriptionUser, error)
	GetUserSubs(
		ctx context.Context,
		limit *api.LimitParam,
		offset *api.OffsetParam,
		subName *api.SubNameParam,
		startDate *api.StartDateParam,
		endDate *api.EndDateParam,
	) ([]*api.SubscriptionUser, error)
	GetUsersForSub(
		ctx context.Context,
		subId api.IdSubParam,
		limit *api.LimitParam,
		offset *api.OffsetParam,
		startDate *api.StartDateParam,
		endDate *api.EndDateParam) ([]*api.SubscriptionUser, error)
	GetSubsForUser(
		ctx context.Context,
		userId *api.IdUserParam,
		limit *api.LimitParam,
		offset *api.OffsetParam,
		subName *api.SubNameParam,
		startDate *api.StartDateParam,
		endDate *api.EndDateParam,
	) ([]*api.SubscriptionUser, error)
}

/*
Paths:
/subscription/id/users/id
*/
type SubscriptionIdUserIdService interface {
	GetUserSubById(ctx context.Context, userId *api.IdUserParam, subId api.IdSubParam) (*api.SubscriptionUser, error)
	UpdateUserSub(ctx context.Context, userId *api.IdUserParam, subId api.IdSubParam, userSub *api.SubscriptionUser) (*api.SubscriptionUser, error)
	DeleteUserSub(ctx context.Context, userId *api.IdUserParam, subId api.IdSubParam) error
	GetUserTotal(ctx context.Context, userId *api.IdUserParam, startDate *api.StartDateParam, endDate *api.EndDateParam) (*api.SubSum, error)
}

type Handler struct {
	api.Unimplemented

	logger             *slog.Logger
	subService         SubscriptionService
	subUserService     SubscriptionUserService
	subIdUserIdService SubscriptionIdUserIdService
}

func NewHandler(
	logger *slog.Logger,
	subService SubscriptionService,
	subUserService SubscriptionUserService,
	subIdUserIdService SubscriptionIdUserIdService,
) *Handler {
	return &Handler{
		logger:             logger,
		subService:         subService,
		subUserService:     subUserService,
		subIdUserIdService: subIdUserIdService,
	}
}
