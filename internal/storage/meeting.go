package storage

import (
	"github.com/racoon-devel/calendar/internal/model"
	"gorm.io/gorm"
)

// MeetingService интерфейс работы со встречами в БД
type MeetingService interface {
	CreateMeeting(m model.Meeting, users []uint) (id uint, err error)
}

func (s service) CreateMeeting(m model.Meeting, users []uint) (id uint, err error) {
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// TODO: invites users
		return tx.Create(&m).Error
	})
	id = m.ID
	return
}
