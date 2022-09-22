package server

import (
	"github.com/racoon-devel/calendar/internal/server/models"
	"github.com/racoon-devel/calendar/internal/server/restapi/operations"
	"github.com/racoon-devel/calendar/internal/server/restapi/operations/meeting"
	"github.com/racoon-devel/calendar/internal/server/restapi/operations/user"
)

func (s *Server) configureAPI(api *operations.ServerAPI) {
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(s.createUser)
	api.UserLoginUserHandler = user.LoginUserHandlerFunc(s.login)
	api.UserGetUsersHandler = user.GetUsersHandlerFunc(s.getUsers)
	api.MeetingCreateMeetingHandler = meeting.CreateMeetingHandlerFunc(s.createMeeting)

	api.KeyAuth = func(accessToken string) (*models.Principal, error) {
		userId, err := s.Calendar.CheckAccessIsGranted(accessToken)
		bigId := models.Principal(userId)
		return &bigId, err
	}
}
