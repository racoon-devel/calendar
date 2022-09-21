package server

import (
	"github.com/racoon-devel/calendar/internal/server/restapi/operations"
	"github.com/racoon-devel/calendar/internal/server/restapi/operations/user"
)

func (s *Server) configureAPI(api *operations.ServerAPI) {
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(s.createUser)
}
