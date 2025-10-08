package handler

import (
	api "SubscriberService/api/generated"
	"SubscriberService/internal/service"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (h *Handler) GetSubscriptionsUsers(w http.ResponseWriter, r *http.Request, params api.GetSubscriptionsUsersParams) {
	const op = "http.subscription_user.GetSubscriptionsUsers"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	limit, offset, subName, startDate, endDate := params.Limit, params.Offset, params.SubName, params.StartDate, params.EndDate

	logger.Debug("Getting subscriptions")
	subs, err := h.subUserService.GetUserSubs(ctx, limit, offset, subName, startDate, endDate)
	if err != nil {
		if errors.Is(err, service.ErrOperationFailed) {
			logger.Error("Failed to get subscriptions with users")
			http.Error(w, "Failed to get subscriptions with users", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrNotfound) {
			logger.Info("User subs not fount")
			http.Error(w, "There is no such subscriptions", http.StatusNotFound)
			return
		}
		logger.Debug(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	logger.Debug("Marshaling JSON")
	jsonData, err := json.Marshal(subs)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	logger.Info("Get subscriptions successful")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) PostSubscriptionsUsers(w http.ResponseWriter, r *http.Request) {
	// вернуть схему для создания
	const op = "http.subscription_user.PostSubscriptionsUsers"
	logger := h.logger.With("op", op)

	logger.Debug("Reading request body")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var subUser api.SubscriptionUser
	logger.Debug("Unmarshalling JSON")
	if err = json.Unmarshal(body, &subUser); err != nil {
		logger.Error("Parsing json error")
		http.Error(w, "Parsing json error", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	logger.Debug("Saving subscription")
	createdSubUser, err := h.subUserService.SaveUserSub(ctx, &subUser)
	if err != nil {
		if errors.Is(err, service.ErrOperationFailed) {
			logger.Error("Failed to save new user sub")
			http.Error(w, "Subscription not saved", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrFailedGetResponseData) {
			logger.Warn("Failed to get response data")
			http.Error(w, "Failed to get response data", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrNotfound) {
			logger.Info("sub not found")
			http.Error(w, "No such user sub", http.StatusNotFound)
			return
		}
		logger.Debug(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	logger.Info("Subscription is saved")
	logger.Debug("Marshaling JSON")
	responseJSON, err := json.Marshal(createdSubUser)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func (h *Handler) GetSubscriptionsUsersUserId(w http.ResponseWriter, r *http.Request, userId api.IdUserParam, params api.GetSubscriptionsUsersUserIdParams) {
	const op = "http.subscription.GetSubscriptionsUsersUserId"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	limit, offset, subName, startDate, endDate := params.Limit, params.Offset, params.SubName, params.StartDate, params.EndDate

	logger.Debug("Getting subscriptions for user")
	data, err := h.subUserService.GetSubsForUser(ctx, &userId, limit, offset, subName, startDate, endDate)
	if err != nil {
		if errors.Is(err, service.ErrOperationFailed) {
			logger.Error("Failed to find data")
			http.Error(w, "Failed to find data", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrNotfound) {
			logger.Info("user sub not found")
			http.Error(w, "No such user sub", http.StatusNotFound)
			return
		}
		logger.Debug(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	logger.Debug("Marshaling JSON")
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	logger.Info("Get subscriptions for user successful")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) GetSubscriptionsSubIdUsers(w http.ResponseWriter, r *http.Request, subId api.IdSubParam, params api.GetSubscriptionsSubIdUsersParams) {
	const op = "http.subscription.GetSubscriptionsUsersUserId"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	limit, offset, startDate, endDate := params.Limit, params.Offset, params.StartDate, params.EndDate

	logger.Debug("Getting subscriptions for user")
	data, err := h.subUserService.GetUsersForSub(ctx, subId, limit, offset, startDate, endDate)
	if err != nil {
		if errors.Is(err, service.ErrOperationFailed) {
			logger.Error("Failed to find data")
			http.Error(w, "Failed to find data", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrNotfound) {
			http.Error(w, "No such user sub", http.StatusNotFound)
			return
		}
		logger.Debug(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	logger.Debug("Marshaling JSON")
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	logger.Info("Get users for subscription successful")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
