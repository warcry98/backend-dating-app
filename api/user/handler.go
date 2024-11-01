package user

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	service *UserService
}

func NewUserHandler(db *gorm.DB) *Handler {
	return &Handler{service: NewUserService(*NewUserRepository(db))}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.service.Register(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		UsernameEmail string `json:"username"`
		Password      string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.service.repo.GetByUsernameOrEmail(credentials.UsernameEmail)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	err = h.service.VerifyPassword(user.Password, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, err := h.service.Login(credentials.UsernameEmail, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "token": token})
}
