package handler

import (
	api "SubscriberService/api/generated"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
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
	UpdateSub(ctx context.Context, sub *api.Subscription) (*api.Subscription, error)
	DeleteSub(ctx context.Context, subId api.IdSubParam) error
}

/*
Paths:
/subscriptions/users
/subscriptions/id/users
/subscriptions/users/id
*/
type SubscriptionUserService interface {
	SaveUserSub(ctx context.Context, userSub *api.SubscriptionUserCreate) (*api.SubscriptionUser, error)
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
		subId *api.IdSubParam,
		limit *api.LimitParam,
		offset *api.OffsetParam,
		startDate *api.StartDateParam,
		endDate *api.EndDateParam) ([]*api.SubscriptionUser, error)
	GetSubsForUser(
		ctx context.Context,
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
	GetUserSubById(ctx context.Context, userId *api.IdUserParam, subId *api.IdSubParam) (*api.SubscriptionUser, error)
	UpdateUserSub(ctx context.Context, userId *api.IdUserParam, subId *api.IdSubParam, userSub *api.SubscriptionUserCreate) (*api.SubscriptionUser, error)
	DeleteUserSub(ctx context.Context, userId *api.IdUserParam, subId *api.IdSubParam) error
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

func (h *Handler) GetSubscriptions(w http.ResponseWriter, r *http.Request, params api.GetSubscriptionsParams) {
	ctx := r.Context()
	limit, offset, subName := params.Limit, params.Offset, params.SubName
	subs, err := h.subService.GetSubs(ctx, limit, offset, subName)
	if err != nil {
		h.logger.Info("Failed to get subscriptions")
		// определяем тип ошибки
		return
	}

	jsonData, err := json.Marshal(subs)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) PostSubscriptions(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var sub api.Subscription

	if err = json.Unmarshal(body, &sub); err != nil {
		http.Error(w, "Parsing json error", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	createdSub, err := h.subService.SaveSub(ctx, &sub)
	if err != nil {
		fmt.Println(err.Error())
		h.logger.Error("Failed to save new Sub")
		return
	}

	responseJSON, err := json.Marshal(createdSub)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}
