package swipe

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Handler struct {
	service SwipeService
}

func NewSwipeHandler(db *gorm.DB) *Handler {
	return &Handler{service: *NewSwipeService(*NewSwipeRepository(db))}
}

func (h *Handler) RecordSwipe(w http.ResponseWriter, r *http.Request) {
	var req SwipeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(req.UserID)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.service.RecordSwipe(userID, req.TargetID, req.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Swipe recorded"})
}
