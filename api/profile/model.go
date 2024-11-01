package profile

type Profile struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	UserID   int    `json:"user_id" gorm:"unique"`
	Fullname string `json:"fullname"`
	Sex      string `json:"sex"` // male or female
	Bio      string `json:"bio"`
	Age      int    `json:"age"`
}
