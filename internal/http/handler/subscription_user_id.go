package handler

import (
	api "SubscriberService/api/generated"
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) GetSubscriptionsUsersUserIdTotal(w http.ResponseWriter, r *http.Request, userId api.IdUserParam, params api.GetSubscriptionsUsersUserIdTotalParams) {
	const op = "http.subscription.GetSubscriptionsUsersUserIdTotal"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	startDate, endDate := params.StartDate, params.EndDate

	logger.Debug("Getting total")
	data, err := h.subIdUserIdService.GetUserTotal(ctx, userId, startDate, endDate)
	if err != nil {
		logger.Error("Failed to get total")
		http.Error(w, "Failed to get total", http.StatusInternalServerError)
		// определяем тип ошибки
		return
	}

	logger.Debug("Marshaling JSON")
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	logger.Info("Get total successful")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) DeleteSubscriptionsSubIdUsersUserId(w http.ResponseWriter, r *http.Request, subId api.IdSubParam, userId api.IdUserParam) {
	const op = "service.subscription.DeleteSubscriptionsSubIdUsersUserId"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	logger.Debug("Trying to delete user sub")
	err := h.subIdUserIdService.DeleteUserSub(ctx, userId, subId)
	if err != nil {
		logger.Error("Failed to delete user sub")
		http.Error(w, "Failed to delete user sub", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetSubscriptionsSubIdUsersUserId(w http.ResponseWriter, r *http.Request, subId api.IdSubParam, userId api.IdUserParam) {
	const op = "service.subscription.GetSubscriptionsSubIdUsersUserId"
	logger := h.logger.With("op", op)
	ctx := r.Context()
	logger.Debug("Trying to get subscription")
	sub, err := h.subIdUserIdService.GetUserSubById(ctx, userId, subId)
	if err != nil {
		logger.Error("Failed to find user subscription")
		http.Error(w, "Failed to find user subscription", http.StatusInternalServerError)
		return
		// определить тип ошибки
	}

	logger.Info("User subscription found")
	logger.Debug("Marshaling JSON")
	responseJSON, err := json.Marshal(sub)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	logger.Info("User sub get successful")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

func (h *Handler) PutSubscriptionsSubIdUsersUserId(w http.ResponseWriter, r *http.Request, subId api.IdSubParam, userId api.IdUserParam) {
	const op = "service.subscription.PutSubscriptionsSubIdUsersUserId"
	logger := h.logger.With("op", op)

	logger.Debug("Reading request body")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var userSub api.SubscriptionUserCreate
	logger.Debug("Unmarshalling JSON")
	if err = json.Unmarshal(body, &userSub); err != nil {
		logger.Error("Parsing json error")
		http.Error(w, "Parsing json error", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	logger.Debug("Updating user subscription")
	updatedSub, err := h.subIdUserIdService.UpdateUserSub(ctx, userId, subId, &userSub)
	if err != nil {
		// определить тип ошибки
		logger.Error("Failed to update user sub")
		http.Error(w, "User subscription not updated", http.StatusInternalServerError)
		return
	}
	logger.Info("User subscription is updated")
	logger.Debug("Marshaling JSON")
	responseJSON, err := json.Marshal(updatedSub)
	if err != nil {
		logger.Error("Failed to marshal JSON")
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	logger.Info("User subscription is updated successful")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)

}
