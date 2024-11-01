package user

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Prefer    string `json:"prefer"` // male or female
	Password  string `json:"password"`
	IsPremium bool   `json:"is_premium" gorm:"default:false"`
	Verified  bool   `json:"verified" gorm:"default:false"`
}
