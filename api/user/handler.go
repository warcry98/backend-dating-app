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

// RegisterHandler handles user register.
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param email body string true "Email"
// @Param prefer body string true "Prefer"
// @Param password body string true "Password"
// @Success 201 {object} map[string]string
// @Failure 400 {object} any
// @Router /auth/register [post]
func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}
	creds := User{
		Username:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Prefer:    r.FormValue("prefer"),
		Password:  r.FormValue("password"),
		IsPremium: false,
		Verified:  false,
	}
	err := h.service.Register(creds)
	if err != nil {
		http.Error(w, `{"error": "Failed to create user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginHandler handles user register.
// @Summary Login user
// @Description Login user in the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 201 {object} map[string]string
// @Failure 400 {object} any
// @Router /auth/login [post]
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	type Credentials struct {
		UsernameEmail string `json:"username"`
		Password      string `json:"password"`
	}
	credentials := Credentials{
		UsernameEmail: r.FormValue("username"),
		Password:      r.FormValue("password"),
	}

	user, err := h.service.repo.GetByUsernameOrEmail(credentials.UsernameEmail)
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusUnauthorized)
		return
	}

	err = h.service.VerifyPassword(user.Password, credentials.Password)
	if err != nil {
		http.Error(w, `{"error": "Invalid password"}`, http.StatusUnauthorized)
		return
	}

	token, err := h.service.Login(credentials.UsernameEmail, credentials.Password)
	if err != nil {
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "token": token})
}
