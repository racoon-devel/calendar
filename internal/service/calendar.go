package service

import (
	"github.com/racoon-devel/calendar/internal/recurrent"
	"github.com/racoon-devel/calendar/internal/storage"
)

type Calendar interface {
	User
	Meeting

	LoadCache() error
}

type calendar struct {
	db storage.Service
	m  meetingCache
}

func NewCalendar(service storage.Service) Calendar {
	return &calendar{
		db: service,
		m: meetingCache{
			userMeetings: make(map[uint]*recurrent.RuleList),
		},
	}
}

func (c *calendar) LoadCache() error {
	return nil
}
