package profile

import "gorm.io/gorm"

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r ProfileRepository) GetProfiles(userID int, prefer string, date string, limit int) ([]Profile, error) {
	var profiles []Profile
	err := r.db.Raw(`
		SELECT * FROM profiles
		WHERE user_id NOT IN (
			SELECT target_id from swipes WHERE user_id = ? AND date = ?
		) AND user_id != ? AND sex = ?
		LIMIT ?`, userID, date, userID, prefer, limit).Scan(&profiles).Error

	return profiles, err
}

func (r ProfileRepository) GetOwnProfile(userID int) (Profile, error) {
	var profile Profile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	return profile, err
}

func (r ProfileRepository) CreateProfile(profile Profile) error {
	return r.db.Create(&profile).Error
}

func (r ProfileRepository) UpdateProfile(updatedProfile Profile) error {
	return r.db.Model(&Profile{}).Where("user_id = ?", updatedProfile.UserID).Updates(updatedProfile).Error
}
