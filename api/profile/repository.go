package profile

import "gorm.io/gorm"

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r ProfileRepository) CreateProfile(profile Profile) error {
	return r.db.Create(&profile).Error
}

func (r ProfileRepository) UpdateProfile(userID int, updatedProfile Profile) error {
	return r.db.Model(&Profile{}).Where("user_id = ?", userID).Updates(updatedProfile).Error
}
