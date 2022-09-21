package server

import (
	"errors"
	"github.com/apex/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/racoon-devel/calendar/internal/model"
	"github.com/racoon-devel/calendar/internal/server/models"
	"github.com/racoon-devel/calendar/internal/server/restapi/operations/user"
	"github.com/racoon-devel/calendar/internal/service"
)

func (s *Server) createUser(params user.CreateUserParams) middleware.Responder {
	logCtx := log.WithFields(&log.Fields{
		"from": "rest",
		"req":  "createUser",
	})

	logCtx.Debug("creating user requested")

	id, err := s.Calendar.CreateUser(model.User{
		Login:        string(*params.User.Login),
		PasswordHash: string(*params.User.Password),
		Name:         string(*params.User.Name),
		Surname:      string(*params.User.Surname),
	})
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			logCtx.Errorf("attempt to add existing user: %s", err)
			errMessage := service.ErrUserAlreadyExists.Error()
			return user.NewCreateUserConflict().WithPayload(&models.CreateUserError{Code: 409, Message: &errMessage})
		}

		logCtx.Errorf("process request failed: %s", err)
		return user.NewCreateUserInternalServerError()
	}

	bigId := int64(id)
	return user.NewCreateUserCreated().WithPayload(&models.CreateUserResponse{ID: &bigId})
}

func (s *Server) login(params user.LoginUserParams) middleware.Responder {
	logCtx := log.WithFields(&log.Fields{
		"from": "rest",
		"req":  "login",
		"user": *params.Credentials.Login,
	})

	logCtx.Debug("login requested")

	accessToken, err := s.Calendar.Login(string(*params.Credentials.Login), string(*params.Credentials.Password))
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			logCtx.Errorf("login failed: %s", err)
			errMessage := service.ErrInvalidCredentials.Error()
			return user.NewLoginUserForbidden().WithPayload(&models.LoginError{Code: 401, Message: &errMessage})
		}

		logCtx.Errorf("process request failed: %s", err)
		return user.NewLoginUserInternalServerError()
	}

	logCtx.Debug("User login successful")

	return user.NewLoginUserOK().WithPayload(&models.LoginResponse{AccessToken: accessToken})
}
