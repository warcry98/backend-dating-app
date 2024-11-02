package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&User{})
	return db
}

func TestRegisterUser(t *testing.T) {
	db := setupTestDB()
	handler := NewUserHandler(db)

	router := http.NewServeMux()
	router.HandleFunc("/auth/register", handler.RegisterUser)

	payload := map[string]string{
		"username": "testuser",
		"email":    "testuser@example.com",
		"prefer":   "both",
		"password": "testpassword",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusCreated, response.Code)
	var respBody map[string]string
	err := json.Unmarshal(response.Body.Bytes(), &respBody)
	assert.Nil(t, err)
	assert.Equal(t, "User registered successfully", respBody["message"])
}

func TestLoginUser(t *testing.T) {
	db := setupTestDB()
	userRepo := NewUserRepository(db)
	userService := NewUserService(*userRepo)
	handler := NewUserHandler(db)

	router := http.NewServeMux()
	router.HandleFunc("/auth/login", handler.LoginUser)

	hashPass, err := hashPassword("test1234")
	assert.Nil(t, err)
	dataUser := User{
		Username:  "test",
		Email:     "test@mail.com",
		Prefer:    "female",
		Password:  hashPass,
		Verified:  false,
		IsPremium: false,
	}
	err = userService.Register(dataUser)
	assert.Nil(t, err)

	payload := map[string]string{
		"username": "test@mail.com",
		"password": "test1234",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	var respBody map[string]string
	err = json.Unmarshal(response.Body.Bytes(), &respBody)
	assert.Nil(t, err)
	assert.Equal(t, "Login successful", respBody["message"])
	assert.NotEmpty(t, respBody["token"])
}
