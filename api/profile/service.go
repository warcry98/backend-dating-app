package profile

type ProfileService struct {
	repo ProfileRepository
}

func NewProfileService(repo ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s ProfileService) GetProfiles(userID int, prefer string, date string, limit int) ([]Profile, error) {
	return s.repo.GetProfiles(userID, prefer, date, limit)
}

func (s ProfileService) GetOwnProfile(userID int) (Profile, error) {
	return s.repo.GetOwnProfile(userID)
}

func (s ProfileService) CreateProfile(profile Profile) error {
	return s.repo.CreateProfile(profile)
}

func (s ProfileService) UpdateProfile(updatedProfile Profile) error {
	return s.repo.UpdateProfile(updatedProfile)
}
