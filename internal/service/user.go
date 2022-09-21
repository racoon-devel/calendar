package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/racoon-devel/calendar/internal/model"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

// User описывает интерфейс сервиса менеджмента пользователей
type User interface {
	CreateUser(user model.User) (id uint, err error)
	Login(login, password string) (accessToken string, err error)
	CheckAccessIsGranted(accessToken string) (id uint, err error)
}

var (
	// ErrUserAlreadyExists ошибка возникающая при попытке добавить уже существующего пользователя
	ErrUserAlreadyExists = errors.New("user with the login already exists")
	// ErrInvalidCredentials неправильно указали логин и пароль
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type userService struct {
	mutex       sync.RWMutex
	loginToUser map[string]*model.User
	idToUser    map[uint]*model.User
}

var authSecretKey = []byte("6sbd92736adnlnjlanlsn3833")

func (c *calendar) loadUsers() error {
	c.u.mutex.Lock()
	defer c.u.mutex.Unlock()

	users, err := c.db.LoadUsers()
	if err != nil {
		return err
	}

	for i := range users {
		c.u.idToUser[users[i].ID] = &users[i]
		c.u.loginToUser[users[i].Login] = &users[i]
	}

	return nil
}

func (s *calendar) CreateUser(user model.User) (id uint, err error) {
	s.u.mutex.Lock()
	defer s.u.mutex.Unlock()

	if _, ok := s.u.loginToUser[user.Login]; ok {
		return 0, ErrUserAlreadyExists
	}

	// хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("cannot generate password hash: %w", err)
		return
	}
	user.PasswordHash = string(hashedPassword)

	id, err = s.db.CreateUser(user)
	if err != nil {
		err = fmt.Errorf("couldn't store user to database: %w", err)
		return
	}

	user.ID = id
	s.u.idToUser[id] = &user
	s.u.loginToUser[user.Login] = &user

	return id, nil
}

func (s *calendar) Login(login, password string) (accessToken string, err error) {
	s.u.mutex.RLock()
	defer s.u.mutex.RUnlock()

	u, ok := s.u.loginToUser[login]
	if !ok {
		err = ErrInvalidCredentials
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		err = fmt.Errorf("%w: password mismatch", ErrInvalidCredentials)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID

	accessToken, err = token.SignedString(authSecretKey)
	if err != nil {
		err = fmt.Errorf("cannot sign jwt: %w", err)
	}
	return
}

type authClaims struct {
	Id uint `json:"id"`
}

func (c authClaims) Valid() error {
	return nil
}

func (s *calendar) CheckAccessIsGranted(accessToken string) (id uint, err error) {
	claims := authClaims{}

	_, err = jwt.ParseWithClaims(
		accessToken,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return authSecretKey, nil
		},
	)

	id = claims.Id
	return
}
