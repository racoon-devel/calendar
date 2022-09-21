package storage

import (
	"fmt"
	"github.com/racoon-devel/calendar/internal/model"
)

// UserService интерфейс для работы с записями в БД о пользователях
type UserService interface {
	LoadUsers() ([]model.User, error)
	CreateUser(user model.User) (id uint, err error)
}

func (s service) LoadUsers() ([]model.User, error) {
	result := make([]model.User, 0)
	if err := s.db.Find(&result).Error; err != nil {
		return nil, fmt.Errorf("fetch users failed: %w", err)
	}

	return result, nil
}

func (s service) CreateUser(user model.User) (id uint, err error) {
	err = s.db.Create(&user).Error
	id = user.ID
	return
}
