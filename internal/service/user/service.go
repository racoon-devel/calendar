package user

import "github.com/racoon-devel/calendar/internal/model"

type Service struct {
}

func (s *Service) CreateUser(user *model.User) (id uint, err error) {
	return 0, nil
}

func (s *Service) Login(login, password string) (success bool, err error) {
	return false, nil
}

func (s *Service) IsAccessGranted(accessToken string) (id uint, ok bool) {
	return 0, false
}
