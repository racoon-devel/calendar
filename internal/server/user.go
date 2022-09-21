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

	id, err := s.Calendar.CreateUser(&model.User{
		Login:        string(*params.User.Login),
		PasswordHash: string(*params.User.Password),
		Name:         string(*params.User.Name),
		Surname:      string(*params.User.Surname),
	})
	if err != nil {
		if errors.Is(err, service.UserAlreadyExists) {
			logCtx.Errorf("attempt to add existing user: %s", err)
			errMessage := service.UserAlreadyExists.Error()
			return user.NewCreateUserConflict().WithPayload(&models.CreateUserError{Code: 409, Message: &errMessage})
		}

		logCtx.Errorf("process request failed: %s", err)
		return user.NewCreateUserInternalServerError()
	}

	bigId := int64(id)
	return user.NewCreateUserCreated().WithPayload(&models.CreateUserResponse{ID: &bigId})
}
