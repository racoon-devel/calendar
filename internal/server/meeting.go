package server

import (
	"errors"
	"github.com/apex/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/racoon-devel/calendar/internal/model"
	"github.com/racoon-devel/calendar/internal/server/models"
	"github.com/racoon-devel/calendar/internal/server/restapi/operations/meeting"
	"github.com/racoon-devel/calendar/internal/service"
	"time"
)

func fixIncomingUserList(inUsers []int64, myself uint) []uint {
	users := make([]uint, 0, len(inUsers))
	dupMap := make(map[uint]bool)
	for i := range inUsers {
		id := uint(inUsers[i])
		if id != myself {
			if _, ok := dupMap[id]; !ok {
				users = append(users, id)
				dupMap[id] = true
			}
		}
	}

	return users
}

func (s *Server) createMeeting(params meeting.CreateMeetingParams, userId *models.Principal) middleware.Responder {
	logCtx := log.WithFields(&log.Fields{
		"from": "rest",
		"req":  "createMeeting",
		"user": *userId,
	})

	m := model.Meeting{
		Owner:       uint(*userId),
		Title:       *params.Meeting.Title,
		Description: params.Meeting.Description,
		Private:     params.Meeting.Private,
		Duration:    time.Duration(*params.Meeting.Duration) * time.Minute,
	}

	if params.Meeting.Rrule != "" {
		m.RRule.Valid = true
		m.RRule.String = params.Meeting.Rrule
	}

	startTime, err := time.Parse(time.RFC3339, *params.Meeting.StartTime)
	if err != nil {
		logCtx.Errorf("cannot parse date string %s: %s", *params.Meeting.StartTime, err)
		code := int64(400)
		message := "invalid date format"
		return meeting.NewCreateMeetingBadRequest().WithPayload(&models.CreateMeetingError{
			Code:    &code,
			Message: &message,
		})
	}

	m.StartTime = startTime
	if params.Meeting.Notify > 0 {
		m.Notify.Valid = true
		m.Notify.Int32 = int32(params.Meeting.Notify)
	}

	// из массива пользователей на всякий случай уберем текущего и повторы
	users := fixIncomingUserList(params.Meeting.Users, uint(*userId))

	id, err := s.Calendar.CreateMeeting(m, users)
	if err != nil {
		code := int64(500)
		message := "Unknown error"
		logCtx.Errorf("cannot create meeting: %s", err)
		if errors.Is(err, service.ErrCannotParseRRule) {
			code = 400
			message = service.ErrCannotParseRRule.Error()
		} else if errors.Is(err, service.ErrUserIsNotExist) {
			code = 404
			message = service.ErrUserIsNotExist.Error()
		}

		if code != 500 {
			return meeting.NewCreateMeetingBadRequest().WithPayload(&models.CreateMeetingError{
				Code:    &code,
				Message: &message,
			})
		}

		return meeting.NewCreateMeetingInternalServerError()
	}
	bigId := int64(id)
	return meeting.NewCreateMeetingCreated().WithPayload(&models.CreateMeetingResponse{ID: &bigId})
}
