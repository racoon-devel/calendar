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
		err := tx.Create(&m).Error
		if err == nil {
			// создаем приглашение для всех пользователей из списка
			for _, u := range users {
				invite := model.Invite{
					UserID:    u,
					MeetingID: m.ID,
					Approved:  false,
				}
				if err = tx.Create(&invite).Error; err != nil {
					return err
				}
			}
		}

		return err
	})
	id = m.ID
	return
}
