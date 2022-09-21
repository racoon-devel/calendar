package storage

import "github.com/racoon-devel/calendar/internal/model"

// UserService интерфейс для работы с записями в БД о пользователях
type UserService interface {
	LoadUsers() ([]model.User, error)
	CreateUser(user *model.User) (id uint, err error)
}

func (s service) LoadUsers() ([]model.User, error) {
	return nil, nil
}

func (s service) CreateUser(user *model.User) (id uint, err error) {
	return 0, nil
}
