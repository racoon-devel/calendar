package service

import "github.com/racoon-devel/calendar/internal/model"

type Meeting interface {
	CreateMeeting(m model.Meeting, users []uint) (id uint, err error)
}

func (c *calendar) CreateMeeting(m model.Meeting, users []uint) (id uint, err error) {
	/**
	0. провалидировать RRULE
	1. проверить, что текущий пользовать сам не занят в указанное время
	2. проверить, что все указанные пользователи существуют
	3. создать встречу
	4. добавить в кеш
	*/
	id, err = c.db.CreateMeeting(m, users)
	return
}
