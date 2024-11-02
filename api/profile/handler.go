package profile

import (
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Handler struct {
	service ProfileService
}

func NewProfileHandler(db *gorm.DB) *Handler {
	return &Handler{service: *NewProfileService(*NewProfileRepository(db))}
}

func (h *Handler) GetProfiles(w http.ResponseWriter, r *http.Request) {
	var profiles []Profile
	currentDate := time.Now().Format("2006-01-02")
	userID := r.Context().Value("userID").(int)
	prefer := r.Context().Value("prefer").(string)

	profiles, err := h.service.GetProfiles(userID, prefer, currentDate, 10)
	if err != nil {
		http.Error(w, "Failed to get profiles", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "Request success", "data": profiles})
}

func (h *Handler) GetOwnProfile(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	userID := r.Context().Value("userID").(int)

	profile, err := h.service.GetOwnProfile(userID)
	if err != nil {
		http.Error(w, "Failed to get profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "Request success", "data": profile})
}

func (h *Handler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID")
	profile.UserID = userID.(int)

	err := h.service.CreateProfile(profile)
	if err != nil {
		http.Error(w, "Failed to create profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile created"})
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID")
	profile.UserID = userID.(int)

	err := h.service.UpdateProfile(profile)
	if err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated"})
}
