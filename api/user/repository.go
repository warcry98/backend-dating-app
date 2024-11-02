package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	NewUserRepository(db *gorm.DB) *UserRepository
	CreateUser(user User) error
	GetByUsernameOrEmail(username_or_email string) (User, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetByUsernameOrEmail(username_or_email string) (User, error) {
	var user User
	err := r.db.Where("username = ? OR email = ?", username_or_email, username_or_email).Find(&user).Error
	return user, err
}
