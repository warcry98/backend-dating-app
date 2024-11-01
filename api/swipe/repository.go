package swipe

import (
	"gorm.io/gorm"
)

type SwipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) *SwipeRepository {
	return &SwipeRepository{db: db}
}

func (r SwipeRepository) CountUserSwipes(userID int, date string) (int, error) {
	var count int64
	err := r.db.Model(&Swipe{}).Where("user_id = ? AND date = ?", userID, date).Count(&count).Error
	return int(count), err
}

func (r SwipeRepository) CreateSwipe(swipe Swipe) error {
	return r.db.Create(&swipe).Error
}
