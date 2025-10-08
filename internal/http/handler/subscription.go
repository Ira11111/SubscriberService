package handler

import (
	api "SubscriberService/api/generated"
	"SubscriberService/internal/service"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (h *Handler) GetSubscriptions(w http.ResponseWriter, r *http.Request, params api.GetSubscriptionsParams) {
	const op = "http.subscription.GetSubscriptions"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	limit, offset, subName := params.Limit, params.Offset, params.SubName

	logger.Debug("Getting subscriptions")
	subs, err := h.subService.GetSubs(ctx, limit, offset, subName)
	if err != nil {
		if errors.Is(err, service.ErrNotfound) {
			logger.Info("no such subs")
			http.Error(w, "data not fount", http.StatusNotFound)
			return
		}
		if errors.Is(err, service.ErrOperationFailed) {
			logger.Error("Failed to get subscriptions")
		}
		http.Error(w, "Failed to get subscriptions", http.StatusInternalServerError)
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

func (h *Handler) PostSubscriptions(w http.ResponseWriter, r *http.Request) {
	const op = "http.subscription.PostSubscription"
	logger := h.logger.With("op", op)

	logger.Debug("Reading request body")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var sub api.SubscriptionCreate
	logger.Debug("Unmarshalling JSON")
	if err = json.Unmarshal(body, &sub); err != nil {
		logger.Error("Parsing json error")
		http.Error(w, "Parsing json error", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	logger.Debug("Saving subscription")
	createdSub, err := h.subService.SaveSub(ctx, &sub)

	if err != nil {
		if errors.Is(err, service.ErrOperationFailed) {
			logger.Error("Failed to save new Sub")
			http.Error(w, "Subscription not saved", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrFailedGetResponseData) {
			logger.Warn("Failed to get response data")
			http.Error(w, "Failed to get response data", http.StatusInternalServerError)
			return
		}
		logger.Debug(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	logger.Info("Subscription is saved")
	logger.Debug("Marshaling JSON")
	responseJSON, err := json.Marshal(createdSub)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func (h *Handler) DeleteSubscriptionsSubId(w http.ResponseWriter, r *http.Request, subId api.IdSubParam) {
	const op = "service.subscription.DeleteSubscriptionsSubId"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	logger.Debug("Trying to delete sub")
	err := h.subService.DeleteSub(ctx, subId)
	if err != nil {
		logger.Error("Failed to delete sub")
		if errors.Is(err, service.ErrOperationFailed) {
			http.Error(w, "Failed to delete sub", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrNotfound) {
			http.Error(w, "No such sub", http.StatusNotFound)
			return
		}
		logger.Debug(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetSubscriptionsSubId(w http.ResponseWriter, r *http.Request, subId api.IdSubParam) {
	const op = "service.subscription.GetSubscriptionsSubId"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	logger.Debug("Trying to get subscription")
	sub, err := h.subService.GetSubById(ctx, subId)
	if err != nil {
		if errors.Is(err, service.ErrOperationFailed) {
			logger.Error("Failed to find subscription")
			http.Error(w, "Failed to find subscription", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrNotfound) {
			http.Error(w, "No such sub", http.StatusNotFound)
			return
		}
		logger.Debug(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	logger.Info("Subscription found")
	logger.Debug("Marshaling JSON")
	responseJSON, err := json.Marshal(sub)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

func (h *Handler) PutSubscriptionsSubId(w http.ResponseWriter, r *http.Request, subId api.IdSubParam) {
	const op = "service.subscription.PutSubscriptionsSubId"
	logger := h.logger.With("op", op)

	logger.Debug("Reading request body")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var sub api.SubscriptionCreate
	logger.Debug("Unmarshalling JSON")
	if err = json.Unmarshal(body, &sub); err != nil {
		logger.Error("Parsing json error")
		http.Error(w, "Parsing json error", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	logger.Debug("Updating subscription")
	createdSub, err := h.subService.UpdateSub(ctx, &sub, subId)
	if err != nil {
		if errors.Is(err, service.ErrOperationFailed) {
			logger.Error("Failed to update subscription")
			http.Error(w, "Failed to update subscription", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, service.ErrNotfound) {
			http.Error(w, "No such sub", http.StatusNotFound)
			return
		}
		logger.Debug(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	logger.Info("Subscription is updated")
	logger.Debug("Marshaling JSON")
	responseJSON, err := json.Marshal(createdSub)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}
