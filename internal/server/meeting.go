package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/racoon-devel/calendar/internal/server/models"
	"github.com/racoon-devel/calendar/internal/server/restapi/operations/meeting"
)

func (s *Server) createMeeting(params meeting.CreateMeetingParams, userId *models.Principal) middleware.Responder {
	return meeting.NewCreateMeetingBadRequest()
}
