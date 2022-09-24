package storage

import (
	"errors"
	"fmt"
	"github.com/racoon-devel/calendar/internal/config"
	"github.com/racoon-devel/calendar/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service общий интерфейс, предоставляющий все возможности работы с БД
type Service interface {
	UserService
	MeetingService
}

type service struct {
	db *gorm.DB
}

var (
	// ErrRecordAlreadyExists ошибка возникает в случае, если пытаемся добавить уже существующую запись
	ErrRecordAlreadyExists = errors.New("record is already exists")
)

// New выполняет подключение к БД и возвращает интерфейс для работы
func New(conParams *config.Database) (srv Service, err error) {
	s := service{}
	srv = &s

	conStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conParams.Host, conParams.User, conParams.Password, conParams.Name, conParams.Port,
	)
	s.db, err = gorm.Open(postgres.Open(conStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}

	err = s.db.AutoMigrate(&model.User{}, &model.Meeting{}, &model.Invite{})
	return
}
