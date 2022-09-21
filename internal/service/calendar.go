package service

import (
	"github.com/racoon-devel/calendar/internal/service/user"
	"github.com/racoon-devel/calendar/internal/storage"
)

type Calendar interface {
	User
}

type calendar struct {
	user.Service
}

func NewCalendar(service storage.Service) Calendar {
	return &calendar{}
}
