package service

import (
	"github.com/racoon-devel/calendar/internal/storage"
)

type Calendar interface {
	User

	LoadCache() error
}

type calendar struct {
	db storage.Service
}

func NewCalendar(service storage.Service) Calendar {
	return &calendar{
		db: service,
	}
}

func (c *calendar) LoadCache() error {
	return nil
}
