package server

import (
	"fmt"
	"github.com/apex/log"
	"github.com/go-openapi/loads"
	"github.com/racoon-devel/calendar/internal/server/restapi"
	"github.com/racoon-devel/calendar/internal/server/restapi/operations"
	"github.com/racoon-devel/calendar/internal/service"
	"net"
	"strconv"
)

// Server имлементирует REST HTTP API сервер, сгенерированный через Swagger
// Для установки обработчиков запросов используется поле Calendar
type Server struct {
	Calendar service.Calendar
	srv      *restapi.Server
}

func (s *Server) ListenAndServer(addr string) error {
	if s.srv == nil {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			return err
		}

		// создаем хендлеры API по умолчанию
		api := operations.NewServerAPI(swaggerSpec)
		if s.Calendar != nil {
			s.configureAPI(api)
		}

		// устанавливаем свой логгер
		logCtx := log.WithField("from", "rest")
		api.Logger = func(s string, i ...interface{}) {
			logCtx.Infof(s, i...)
		}

		// создаем и настраиваем сервер
		s.srv = restapi.NewServer(api)
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("cannot parse addr: %w", err)
	}

	portNum, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return fmt.Errorf("parse port failed: %w", err)
	}

	s.srv.Host = host
	s.srv.Port = int(portNum)

	if err := s.srv.Listen(); err != nil {
		return fmt.Errorf("cannot start server: %w", err)
	}

	return s.srv.Serve()
}

func (s *Server) Shutdown() error {
	if s.srv != nil {
		return s.srv.Shutdown()
	}

	return nil
}
