package swipe

import "time"

type Swipe struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"index"`
	TargetID  int       `json:"target_id" gorm:"index"`
	Action    string    `json:"action"` // "like" or "pass"
	Date      string    `json:"date" gorm:"type:date"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type SwipeRequest struct {
	UserID   string `json:"user_id"`
	TargetID int    `json:"target_id"`
	Action   string `json:"action"` // "like" or "pass"
}
