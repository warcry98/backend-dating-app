package user

import "errors"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user User) error {
	return s.repo.Create(user)
}

func (s *UserService) Login(username_or_email, password string) (User, error) {
	user, err := s.repo.GetByUsernameOrEmail(username_or_email)
	if err != nil {
		return User{}, err
	}
	if user.Password != password {
		return User{}, errors.New("invalid credentials")
	}
	return user, nil
}
