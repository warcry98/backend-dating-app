package user

import (
	"backend-dating-app/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

type UserServiceInterface interface {
	Register(user User) error
	Login(username_or_email, password string) (string, error)
	VerifyPassword(hashedPassword, plainPassword string) error
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s *UserService) Login(username_or_email, password string) (string, error) {
	user, err := s.repo.GetByUsernameOrEmail(username_or_email)
	if err != nil {
		return "", err
	}

	err = s.VerifyPassword(user.Password, password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s UserService) VerifyPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GenerateJWT(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"prefer":    user.Prefer,
		"isPremium": user.IsPremium,
		"verified":  user.Verified,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(config.LoadConfig().SecretKey))
}
