package service

import (
	"errors"
	"fmt"
	"github.com/racoon-devel/calendar/internal/model"
	"github.com/racoon-devel/calendar/internal/recurrent"
	"sync"
	"time"
)

type Meeting interface {
	CreateMeeting(m model.Meeting, users []uint) (id uint, err error)
}

var (
	// ErrCannotParseRRule не удалось распарсить рекуррентное выражение
	ErrCannotParseRRule = errors.New("cannot parse RRule")

	// ErrUserIsNotExist приглашенный пользователь не существует
	ErrUserIsNotExist = errors.New("invited user is not exist")
)

type meetingCache struct {
	mutex        sync.RWMutex
	userMeetings map[uint]*recurrent.RuleList
}

func (c *meetingCache) append(userId uint, r recurrent.Rule) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	list, ok := c.userMeetings[userId]
	if !ok {
		list = &recurrent.RuleList{}
		list.Add(r)
		c.userMeetings[userId] = list
		return
	}
	list.Add(r)
}

func (c *meetingCache) isUserBusy(userId uint, startTime time.Time, dur time.Duration) bool {
	list, ok := c.userMeetings[userId]
	if !ok {
		return false
	}

	return list.IsIntersect(startTime, dur)
}

func (c *calendar) CreateMeeting(m model.Meeting, users []uint) (id uint, err error) {
	var r recurrent.Rule

	// 1. парсим и валидируем RRule
	if m.RRule.Valid {
		r, err = recurrent.Parse(m.RRule.String, m.StartTime, m.Duration)
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrCannotParseRRule, err)
			return
		}
	} else {
		r = recurrent.Once(m.StartTime, m.Duration)
	}

	// 2. проверяем, что пользователи существуют. Так как нету удаления пользователей, необязательно делать это атомарно
	var exist bool
	exist, err = c.db.UsersExist(users)
	if err != nil {
		err = fmt.Errorf("cannot check that all users are exist: %w", err)
		return
	}
	if !exist {
		err = ErrUserIsNotExist
		return
	}

	// казалось бы логично проверить, что текущий пользователь не занят на время задачи, однако
	// Google календарь так не делает

	// 3. создаем встречу
	id, err = c.db.CreateMeeting(m, users)

	// 4. добавляем в кеш
	if err != nil {
		c.m.append(m.Owner, r)
	}
	return
}
