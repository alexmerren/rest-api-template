package rest

import (
	"context"
	"fmt"
	"net/http"
	"rest-api-template/internal/domain/entities"
	"rest-api-template/internal/domain/repositories"
)

type RESTServer struct {
	usecases   repositories.ContactUseCases
	logger     repositories.Logger
	httpServer *http.Server
}

func NewRESTServer(
	usecases repositories.ContactUseCases,
	logger repositories.Logger,
	port int,
) *RESTServer {
	return &RESTServer{
		usecases: usecases,
		logger:   logger,
		httpServer: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
	}
}

func (s *RESTServer) Start() error {
	s.logger.Info(fmt.Sprintf("Starting HTTP server : %s", s.httpServer.Addr))
	if err := s.mapRoutes(); err != nil {
		s.logger.Error(err)
		return entities.NewInternalError("could not map routes", err)
	}
	if err := s.httpServer.ListenAndServe(); err != nil {
		s.logger.Error(err)
		return entities.NewInternalError("there was an error starting http server", err)
	}
	return nil
}

func (s *RESTServer) Stop(ctx context.Context) error {
	s.logger.Info("Stopping HTTP server")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Info("server shutdown failed")
		return err
	}
	return nil
}
