package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/racoon-devel/calendar/internal/model"
	"github.com/racoon-devel/calendar/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

// User описывает интерфейс сервиса менеджмента пользователей
type User interface {
	CreateUser(user model.User) (id uint, err error)
	Login(login, password string) (accessToken string, err error)
	CheckAccessIsGranted(accessToken string) (id uint, err error)
	GetAllUsers() ([]model.User, error)
}

var (
	// ErrUserAlreadyExists ошибка возникающая при попытке добавить уже существующего пользователя
	ErrUserAlreadyExists = errors.New("user with the login already exists")
	// ErrInvalidCredentials неправильно указали логин и пароль
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// todo: вынести в Env
var authSecretKey = []byte("6sbd92736adnlnjlanlsn3833")

func (c *calendar) CreateUser(user model.User) (id uint, err error) {
	// хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("cannot generate password hash: %w", err)
		return
	}
	user.PasswordHash = string(hashedPassword)

	id, err = c.db.CreateUser(user)
	if err != nil {
		if errors.Is(err, storage.ErrRecordAlreadyExists) {
			err = ErrUserAlreadyExists
			return
		}
		err = fmt.Errorf("couldn't store user to database: %w", err)
		return
	}

	return id, nil
}

func (c *calendar) Login(login, password string) (accessToken string, err error) {
	var u *model.User

	u, err = c.db.FindUserByLogin(login)
	if err != nil {
		err = fmt.Errorf("%w: user not found: %s", ErrInvalidCredentials, err)
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

func (c *calendar) CheckAccessIsGranted(accessToken string) (id uint, err error) {
	claims := authClaims{}

	// Доверяем ИДу юзера из подписанной строчки, в БД лезть незачем с учетом того, что пользователи не удалются
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

func (c *calendar) GetAllUsers() ([]model.User, error) {
	return c.db.LoadUsers()
}
