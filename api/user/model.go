package user

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	IsPremium bool   `json:"is_premium"`
	Verified  bool   `json:"verified"`
}
