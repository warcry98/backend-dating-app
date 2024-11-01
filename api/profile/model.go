package profile

type Profile struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	UserID int    `json:"user_id" gorm:"index"`
	Bio    string `json:"bio"`
	Age    int    `json:"age"`
}
