package service

import (
	"errors"
	"github.com/racoon-devel/calendar/internal/model"
)

// User описывает интерфейс сервиса менеджмента пользователей
type User interface {
	CreateUser(user *model.User) (id uint, err error)
	Login(login, password string) (success bool, err error)
	IsAccessGranted(accessToken string) (id uint, ok bool)
}

var (
	// UserAlreadyExists ошибка возникающая при попытке добавить уже существующего пользователя
	UserAlreadyExists = errors.New("user with the login already exists")
)
