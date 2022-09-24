package storage

import (
	"errors"
	"fmt"
	"github.com/racoon-devel/calendar/internal/model"
	"gorm.io/gorm"
)

// UserService интерфейс для работы с записями в БД о пользователях
type UserService interface {
	LoadUsers() ([]model.User, error)
	CreateUser(user model.User) (id uint, err error)
	FindUserByLogin(login string) (user *model.User, err error)
	UsersExist(users []uint) (bool, error)
}

func (s service) LoadUsers() ([]model.User, error) {
	result := make([]model.User, 0)
	if err := s.db.Find(&result).Error; err != nil {
		return nil, fmt.Errorf("fetch users failed: %w", err)
	}

	return result, nil
}

func (s service) CreateUser(user model.User) (id uint, err error) {
	err = s.db.Transaction(func(tx *gorm.DB) error {
		var u model.User
		if !errors.Is(s.db.Where("login = ?", user.Login).Take(&u).Error, gorm.ErrRecordNotFound) {
			return ErrRecordAlreadyExists
		}

		return s.db.Create(&user).Error
	})

	id = user.ID
	return
}

func (s service) FindUserByLogin(login string) (user *model.User, err error) {
	user = &model.User{}
	err = s.db.Where("login = ?", login).Take(user).Error
	return
}

func (s service) UsersExist(users []uint) (bool, error) {
	if len(users) == 0 {
		return true, nil
	}
	var count int64
	err := s.db.Model(&model.User{}).Where(users).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == int64(len(users)), nil
}
