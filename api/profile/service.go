package profile

type ProfileService struct {
	repo ProfileRepository
}

func (s ProfileService) CreateProfile(profile Profile) error {
	return s.repo.CreateProfile(profile)
}

func (s ProfileService) UpdateProfile(userID int, updatedProfile Profile) error {
	return s.repo.UpdateProfile(userID, updatedProfile)
}
