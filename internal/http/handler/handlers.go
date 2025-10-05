package handler

import (
	api "SubscriberService/api/generated"
	"log/slog"

	"github.com/go-chi/chi/v5"
)

/*
Paths:
/subscriptions
/subscriptions/id
*/
type SubscriptionService interface {
	SaveSub(ctx *chi.Context, sub *api.Subscription) (*api.Subscription, error)
	GetSubs(ctx *chi.Context, limit *api.LimitParam, offset *api.OffsetParam, subName *api.SubNameParam) ([]*api.Subscription, error)
	GetSubById(ctx *chi.Context, subId *api.IdSubParam) (*api.Subscription, error)
	UpdateSub(ctx *chi.Context, sub *api.Subscription) (*api.Subscription, error)
	DeleteSub(ctx *chi.Context, subId *api.IdSubParam) error
}

/*
Paths:
/subscriptions/users
/subscriptions/id/users
/subscriptions/users/id
*/
type SubscriptionUserService interface {
	SaveUserSub(ctx *chi.Context, userSub *api.SubscriptionUserCreate) (*api.SubscriptionUser, error)
	GetUserSubs(
		ctx *chi.Context,
		limit *api.LimitParam,
		offset *api.OffsetParam,
		subName *api.SubNameParam,
		startDate *api.StartDateParam,
		endDate *api.EndDateParam,
	) ([]*api.SubscriptionUser, error)
	GetUsersForSub(
		ctx *chi.Context,
		subId *api.IdSubParam,
		limit *api.LimitParam,
		offset *api.OffsetParam,
		startDate *api.StartDateParam,
		endDate *api.EndDateParam) ([]*api.SubscriptionUser, error)
	GetSubsForUser(
		ctx *chi.Context,
		userid *api.IdUserParam,
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
	GetUserSubById(ctx *chi.Context, userId *api.IdUserParam, subId *api.IdSubParam) (*api.SubscriptionUser, error)
	UpdateUserSub(ctx *chi.Context, userId *api.IdUserParam, subId *api.IdSubParam, userSub *api.SubscriptionUserCreate) (*api.SubscriptionUser, error)
	DeleteUserSub(ctx *chi.Context, userId *api.IdUserParam, subId *api.IdSubParam) error
	GetUserTotal(ctx *chi.Context, userId *api.IdUserParam, startDate *api.StartDateParam, endDate *api.EndDateParam) (*api.SubSum, error)
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
