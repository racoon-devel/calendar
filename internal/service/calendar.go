package service

import (
	"fmt"
	"github.com/racoon-devel/calendar/internal/model"
	"github.com/racoon-devel/calendar/internal/storage"
)

type Calendar interface {
	User

	LoadCache() error
}

type calendar struct {
	db storage.Service
	u  userService
}

func NewCalendar(service storage.Service) Calendar {
	return &calendar{
		db: service,
		u: userService{
			idToUser:    map[uint]*model.User{},
			loginToUser: map[string]*model.User{},
		},
	}
}

func (c *calendar) LoadCache() error {
	if err := c.loadUsers(); err != nil {
		return fmt.Errorf("load all users failed: %w", err)
	}

	return nil
}
