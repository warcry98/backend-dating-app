package swipe

import (
	"errors"
	"time"
)

type SwipeService struct {
	repo SwipeRepository
}

func NewSwipeService(repo SwipeRepository) *SwipeService {
	return &SwipeService{repo: repo}
}

func (s SwipeService) RecordSwipe(userID, targetID int, action string) error {
	currentDate := time.Now().Format("2006-01-02")

	dailyCount, err := s.repo.CountUserSwipes(userID, currentDate)
	if err != nil {
		return err
	}

	if dailyCount >= 10 {
		return errors.New("daily swipe limit reached")
	}

	return s.repo.CreateSwipe(Swipe{
		UserID:   userID,
		TargetID: targetID,
		Action:   action,
		Date:     currentDate,
	})
}
