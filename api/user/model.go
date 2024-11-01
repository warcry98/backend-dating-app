package user

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Prefer    string `json:"prefer" gorm:"type:enum('male', 'female')"`
	Password  string `json:"password"`
	IsPremium bool   `json:"is_premium"`
	Verified  bool   `json:"verified"`
}
